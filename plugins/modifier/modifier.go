package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
)

// ModifierRegisterer is the symbol the plugin loader will be looking for. It must
// implement the plugin.Registerer interface
// https://github.com/luraproject/lura/blob/master/proxy/plugin/modifier.go#L71
var (
	pluginName            = "my-modifier"
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
	fmt.Println(string(r), "registered!!!")
}

var unknownTypeErr = errors.New("unknown request type")

func (r registerer) request(
	extra map[string]interface{},
) func(interface{}) (interface{}, error) {
	// check the cfg. If the modifier requires some configuration,
	// it should be under the name of the plugin.
	// ex: if this modifier required some A and B config params
	/*
	   "extra_config":{
	       "plugin/req-resp-modifier":{
	           "name":["my-modifier-request"],
	           "my-modifier-request":{
	               "option":"value"
	           }
	       }
	   }
	*/
	config, err := parseConfig("request", extra)
	if err != nil {
		logger.Debug(fmt.Sprintf("Cannot initialize configuration: %s", err.Error()))
		return emptyModifier
	}

	logger.Debug(fmt.Sprintf("The plugin %s-request is now running %s", pluginName, config.Opt))
	// return the modifier
	return func(input interface{}) (interface{}, error) {
		req, ok := input.(RequestWrapper)
		if !ok {
			return nil, unknownTypeErr
		}

		fmt.Println("params:", req.Params())
		fmt.Println("headers:", req.Headers())
		fmt.Println("method:", req.Method())
		fmt.Println("url:", req.URL())
		fmt.Println("query:", req.Query())
		fmt.Println("path:", req.Path())

		// Return input directly if you don't modify the response
		return input, nil
	}
}

func (r registerer) response(
	extra map[string]interface{},
) func(interface{}) (interface{}, error) {
	// check the cfg. If the modifier requires some configuration,
	// it should be under the name of the plugin.
	// ex: if this modifier required some A and B config params
	/*
	   "extra_config":{
	       "plugin/req-resp-modifier":{
	           "name":["my-modifier-response"],
	           "my-modifier-response":{
	               "option":"value"
	           }
	       }
	   }
	*/
	config, err := parseConfig("response", extra)
	if err != nil {
		logger.Debug(fmt.Sprintf("Cannot initialize configuration: %s", err.Error()))
		return emptyModifier
	}

	logger.Debug(fmt.Sprintf("The plugin %s-response is now running %s", pluginName, config.Opt))
	// return the modifier
	return func(input interface{}) (interface{}, error) {
		resp, ok := input.(ResponseWrapper)
		if !ok {
			return nil, unknownTypeErr
		}

		logger.Debug("status code:", resp.StatusCode())

		// Return input directly if you don't modify the response
		return input, nil
	}
}

type config struct {
	Opt string `json:"option"`
}

// parseConfig parses the configuration marshaling and unmarshaling into a struct.
// you can also manually check for fields in the extra conffig map
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
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Logger loaded", pluginName))
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
