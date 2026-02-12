package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

var MiddlewareRegisterer = registerer("quota-control-mw")

var (
	logger                Logger          = nil
	ctx                   context.Context = context.Background()
	errRegistererNotFound                 = fmt.Errorf("%s plugin disabled: config not found", MiddlewareRegisterer)
	errUnkownRequestType                  = errors.New("unknown request type")
	errUnkownResponseType                 = errors.New("unknown response type")
)

type registerer string

func (r registerer) RegisterMiddlewares(f func(
	name string,
	middlewareFactory func(map[string]interface{}, func(context.Context, interface{}) (interface{}, error)) func(context.Context, interface{}) (interface{}, error),
),
) {
	f(string(r), r.middlewareFactory)
}

func (r registerer) middlewareFactory(cfg map[string]interface{}, next func(context.Context, interface{}) (interface{}, error)) func(context.Context, interface{}) (interface{}, error) {
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Middleware injected", r))

	// seeking the quota processor, it must be declared in the configuration at service level
	p, err := quotaSelector("llm_control")
	if err != nil {
		logger.Error(fmt.Sprintf("[PLUGIN: %s] Unable to find quota processor", r))
		return next
	}

	quotaProcessor, ok := p.(QuotaProcessor)
	if !ok {
		logger.Error(fmt.Sprintf("[PLUGIN: %s] Unable to load quota processor", r))
		return next
	}

	// activating rules in the quota processor
	// this associates a rule name defined at service level config (gold, silver...) with an arbitrary value, which can get extracted from the request (admin, dev...)
	errs := quotaProcessor.WithRules(
		[2]string{"gold", "admin"},
		[2]string{"silver", "developer"},
		[2]string{"bronze", "management"},
	)
	if errs != nil {
		quotaProcessor.Close()
		return next
	}

	usagePattern, _ := regexp.Compile(`{"prompt_tokens":(\d+),"completion_tokens":(\d+),"total_tokens":(\d+)}`)

	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqw, ok := req.(RequestWrapper)
		if !ok {
			return nil, errUnkownRequestType
		}

		n := time.Now()

		// these values can be extracted from the request via headers, params, JWT claims...
		tier := "admin"
		userId := "1234"

		// weightless pre-check, this will check if the current request has already reached quota limits and should be rejected
		preCheck, _ := quotaProcessor.Process(tier, userId, n, 0)
		s, ok := preCheck.(processResponseWrapper)
		if !ok {
			return nil, errors.New("unexpected error processing quotas")
		}
		if !s.Allowed() {
			return nil, errors.New("quota exceeded")
		}

		resp, err := next(ctx, reqw)
		if err != nil {
			return nil, err
		}
		respw, ok := resp.(ResponseWrapper)
		if !ok {
			return nil, errUnkownResponseType
		}

		// this is the most relevant part: instead of returning a plain response, we wrap around a custom Reader
		// that can analyze each stream chunk
		// check below for its implementation
		return streamResponseWrapper{
			ResponseWrapper: respw,
			stream: streamReader{
				Reader:       respw.Io(),
				usagePattern: usagePattern,
				increaseHandler: func(tokens int64) {
					quotaProcessor.Process(tier, userId, n, tokens)
				},
			},
		}, err
	}
}

func (r registerer) RegisterLogger(in interface{}) {
	l, ok := in.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Logger loaded", r))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

func main() {}

type RequestWrapper interface {
	Context() context.Context
	Params() map[string]string
	Headers() map[string][]string
	Body() io.ReadCloser
	Method() string
	URL() *url.URL
	Query() url.Values
	Path() string
}

type ResponseWrapper interface {
	Context() context.Context
	Request() interface{}
	Data() map[string]interface{}
	IsComplete() bool
	Headers() map[string][]string
	StatusCode() int
	Io() io.Reader
}

type streamResponseWrapper struct {
	ResponseWrapper
	stream io.Reader
}

func (s streamResponseWrapper) Io() io.Reader {
	return s.stream
}

type streamReader struct {
	io.Reader
	usagePattern    *regexp.Regexp
	processor       QuotaProcessor
	increaseHandler func(int64)
}

func (s streamReader) Read(p []byte) (n int, err error) {
	n, err = s.Reader.Read(p)
	if n > 0 {
		data := p[:n]
		matches := s.usagePattern.FindAllSubmatch(data, -1)
		if matches != nil {
			m := matches[0]
			if len(m) == 4 {
				totalTokens, _ := strconv.Atoi(string(m[3]))
				s.increaseHandler(int64(totalTokens))
			}
		}
	}
	return
}

type QuotaProcessor interface {
	Close()
	WithRules(rules ...[2]string) error
	Process(tierValue string, key string, when time.Time, weight int64) (interface{}, error)
}

var quotaSelector func(string) (interface{}, error)

func (r registerer) RegisterQuotaProcessorSelector(s func(string) (interface{}, error)) {
	quotaSelector = s
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Quota processor selector loaded", r))
}

type processResponseWrapper interface {
	Allowed() bool
	BlockedFor() time.Duration
	Details() []interface{}
}

type windowStatusWrapper interface {
	Remaining() int64
	Consumed() int64
	NextRefill() time.Duration
	LevelValue() string
	WindowName() string
	Allowed() bool
}
