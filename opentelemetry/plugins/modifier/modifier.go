package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/krakend/krakend-otel/state"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ModifierRegisterer is the symbol the plugin loader will be looking for. It must
// implement the plugin.Registerer interface
// https://github.com/luraproject/lura/blob/master/proxy/plugin/modifier.go#L71
var (
	pluginName            = "otel-modifier"
	ModifierRegisterer    = registerer(pluginName)
	errRegistererNotFound = fmt.Errorf("%s plugin disabled: config not found", pluginName)
)

type registerer string

// RegisterModifiers is the function the plugin loader will call to register the
// modifier(s) contained in the plugin using the function passed as argument.
// f will register the factoryFunc under the name and mark it as a request
// and/or response modifier.
func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r)+"-request", r.request, true, false)
	f(string(r)+"-response", r.response, false, true)
}

var unknownTypeErr = errors.New("unknown request type")

func (r registerer) request(
	extra map[string]interface{},
) func(interface{}) (interface{}, error) {
	prefix := fmt.Sprintf("[PLUGIN: %s-request]", pluginName)
	config, err := parseConfig("request", extra)
	if err != nil {
		logger.Warning(fmt.Sprintf("%s Cannot initialize configuration: %s",
			prefix, err.Error()))
		return emptyModifier
	}

	logger.Debug(fmt.Sprintf("%s is now running: %s", prefix, config.Opt))
	// return the modifier

	otelConf := state.GlobalConfig()
	otelState := otelConf.OTEL() // get the global configured settings
	tracer := otelState.Tracer()

	return func(input interface{}) (interface{}, error) {
		req, ok := input.(RequestWrapper)
		if !ok {
			logger.Warning(fmt.Sprintf("%s cannot cast to RequestWrapper %q",
				prefix, config.Opt))
			return nil, unknownTypeErr
		}

		// set attributes in the current Span
		logger.Info(fmt.Sprintf("%s setting 'plugin_req_mod_option' = %q",
			prefix, config.Opt))
		rCtx := req.Context()
		span := trace.SpanFromContext(rCtx)
		span.SetAttributes(
			attribute.Key("plugin_req_mod_option").String(config.Opt),
			attribute.Key("plugin_req_mod_query").String(req.Query().Encode()))

		// create a new Span
		spanName := "req_mod_plugin_span"
		_, newSpan := tracer.Start(rCtx, spanName)
		// DO Something here:
		time.Sleep(25 * time.Millisecond)
		newSpan.End()

		// Return input directly if you don't modify the response
		return input, nil
	}
}

func (r registerer) response(
	extra map[string]interface{},
) func(interface{}) (interface{}, error) {

	prefix := fmt.Sprintf("[PLUGIN: %s-response]", pluginName)
	config, err := parseConfig("response", extra)
	if err != nil {
		logger.Warning(fmt.Sprintf("%s Cannot initialize configuration: %s",
			prefix, err.Error()))
		return emptyModifier
	}

	// WARNING: One limitation is that since plugins only receive the extra_config
	// section, we cannot access the endpoint / backend config, and thus apply
	// any overriden opentelemetry options that might be configured.

	logger.Warning(fmt.Sprintf("%s now running %s", prefix, config.Opt))

	// Get Access to the Global Open Telemetry config:
	otelConf := state.GlobalConfig()
	otelState := otelConf.OTEL() // get the global configured state
	tracer := otelState.Tracer()

	// return the modifier
	return func(input interface{}) (interface{}, error) {
		resp, ok := input.(ResponseWrapper)
		if !ok {
			logger.Warning(fmt.Sprintf("%s cannot cast to ResponseWrapper %q",
				prefix, config.Opt))
			return nil, unknownTypeErr
		}

		logger.Info(fmt.Sprintf("%s setting 'plugin_resp_mod_option' = %q",
			prefix, config.Opt))

		// set attributes in the current Span
		rCtx := resp.Context()
		span := trace.SpanFromContext(rCtx)
		span.SetAttributes(attribute.Key("plugin_resp_mod_option").String(config.Opt))

		// create a new Span
		spanName := "resp_mod_plugin_span"
		_, newSpan := tracer.Start(rCtx, spanName)
		// DO Something here:
		time.Sleep(25 * time.Millisecond)
		newSpan.End()

		// Return input directly if you don't modify the response
		return input, nil
	}
}

type config struct {
	Opt string `json:"option"`
}

// parseConfig parses the configuration marshaling and unmarshaling into a struct.
// you can also manually check for fields in the extra config map
func parseConfig(namespace string, extra map[string]interface{}) (*config, error) {
	ns := fmt.Sprintf("%s-%s", pluginName, namespace)
	cfgRaw, ok := extra[ns].(map[string]interface{})
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

func emptyModifier(input interface{}) (interface{}, error) {
	return input, nil
}

// RequestWrapper is an interface for passing proxy request between the krakend pipe
// and the loaded plugins
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

type requestWrapper struct {
	ctx     context.Context
	method  string
	url     *url.URL
	query   url.Values
	path    string
	body    io.ReadCloser
	params  map[string]string
	headers map[string][]string
}

func (r requestWrapper) Context() context.Context     { return r.ctx }
func (r requestWrapper) Method() string               { return r.method }
func (r requestWrapper) URL() *url.URL                { return r.url }
func (r requestWrapper) Query() url.Values            { return r.query }
func (r requestWrapper) Path() string                 { return r.path }
func (r requestWrapper) Body() io.ReadCloser          { return r.body }
func (r requestWrapper) Params() map[string]string    { return r.params }
func (r requestWrapper) Headers() map[string][]string { return r.headers }

// ResponseWrapper is an interface for passing proxy response between the krakend pipe
// and the loaded plugins
type ResponseWrapper interface {
	Context() context.Context
	Data() map[string]interface{}
	Io() io.Reader
	IsComplete() bool
	StatusCode() int
	Headers() map[string][]string
}

type metadataWrapper struct {
	headers    map[string][]string
	statusCode int
}

func (m metadataWrapper) Headers() map[string][]string { return m.headers }
func (m metadataWrapper) StatusCode() int              { return m.statusCode }

type responseWrapper struct {
	ctx        context.Context
	request    interface{}
	data       map[string]interface{}
	isComplete bool
	metadata   metadataWrapper
	io         io.Reader
}

func (r responseWrapper) Context() context.Context     { return r.ctx }
func (r responseWrapper) Request() interface{}         { return r.request }
func (r responseWrapper) Data() map[string]interface{} { return r.data }
func (r responseWrapper) IsComplete() bool             { return r.isComplete }
func (r responseWrapper) Io() io.Reader                { return r.io }
func (r responseWrapper) Headers() map[string][]string { return r.metadata.headers }
func (r responseWrapper) StatusCode() int              { return r.metadata.statusCode }

// This logger is replaced by the RegisterLogger method to load the one from KrakenD
var logger Logger = noopLogger{}

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] modifier plugin loaded", pluginName))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

// Empty logger implementation
type noopLogger struct{}

func (n noopLogger) Debug(_ ...interface{})    {}
func (n noopLogger) Info(_ ...interface{})     {}
func (n noopLogger) Warning(_ ...interface{})  {}
func (n noopLogger) Error(_ ...interface{})    {}
func (n noopLogger) Critical(_ ...interface{}) {}
func (n noopLogger) Fatal(_ ...interface{})    {}

func main() {}
