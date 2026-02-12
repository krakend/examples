package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"

	"github.com/redis/go-redis/v9"
)

var MiddlewareRegisterer = registerer("counter-example-mw")

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
	config, err := r.parseConfig(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("[PLUGIN: %s] Cannot initialize configuration: %s", r, err))
		return nil
	}

	redisClient := redisClientSelector("local_instance")
	if redisClient == nil {
		logger.Error(fmt.Sprintf("[PLUGIN: %s] Cannot find Redis client", r))
	}

	logger.Info(fmt.Sprintf("[PLUGIN: %s] Middleware injected %+v", r, config))
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqw, ok := req.(RequestWrapper)
		if !ok {
			return nil, errUnkownRequestType
		}

		resp, err := next(ctx, reqw)
		if err != nil {
			return nil, err
		}

		respw, ok := resp.(ResponseWrapper)
		if !ok {
			return nil, errUnkownResponseType
		}

		key := fmt.Sprintf("%s:%s", config.KeyPrefix, "some-id")
		redisClient.Incr(reqw.Context(), key)

		return respw, err
	}
}

type config struct {
	KeyPrefix string `json:"key_prefix"`
}

func (r registerer) parseConfig(extra map[string]interface{}) (*config, error) {
	cfgRaw, ok := extra[string(r)].(map[string]interface{})
	if !ok {
		return nil, errRegistererNotFound
	}

	cfg := config{}
	b, err := json.Marshal(cfgRaw)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

var (
	redisClientSelector        func(string) *redis.Client
	redisClusterClientSelector func(context.Context, string, string) *redis.Client
)

func (r registerer) RegisterRedisSelectors(cs func(string) *redis.Client, ccs func(context.Context, string, string) *redis.Client) {
	redisClientSelector = cs
	redisClusterClientSelector = ccs
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Redis selectors loaded", r))
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
