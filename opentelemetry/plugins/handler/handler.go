// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krakend/krakend-otel/state"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	// pluginName is the plugin name
	pluginName = "otel-handler"
	// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
	HandlerRegisterer     = registerer(pluginName)
	errRegistererNotFound = fmt.Errorf("%s plugin disabled: config not found", pluginName)
)

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	// If the plugin requires some configuration, it should be under the name of the plugin. E.g.:
	/*
	   "extra_config":{
	       "plugin/http-server":{
	           "name":["my-handler-plugin"],
	           "my-handler-plugin":{
	               "someOption": "some-value"
	           }
	       }
	   }
	*/
	// The config variable contains all the keys you have defined in the configuration
	// if the key doesn't exists or is not a map the plugin returns an error and the default handler
	config, err := parseConfig(extra)
	if err != nil {
		logger.Warning(fmt.Sprintf("[PLUGIN: %s] cannot parse config %s -> %#v",
			pluginName, err.Error(), extra))
		return h, errRegistererNotFound
	}

	otelConf := state.GlobalConfig()
	otelState := otelConf.OTEL() // get the global configured settings

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tracer := otelState.Tracer()
		rCtx := req.Context()
		span := trace.SpanFromContext(rCtx)

		// TODO: try adding attributes to current span
		span.SetAttributes(attribute.Key("my_plugin_key").String("my_plugin_value"))
		spanName := "my_plugin_name"

		logger.Warning("starting new span")
		ctx, newSpan := tracer.Start(rCtx, spanName)
		req = req.WithContext(ctx)
		h.ServeHTTP(w, req)
		newSpan.End()
		fmt.Printf("DBG: ending span %s", req.URL.String())
		logger.Warning("ending span", req.URL.String(), config.Opt)

		//logger.Warning("config value option: %s", config.Opt)
	}), nil
}

type config struct {
	Opt string `json:"option"`
}

// parseConfig parses the configuration marshaling and unmarshaling into a struct.
// you can also manually check for fields in the extra conffig map
func parseConfig(extra map[string]interface{}) (*config, error) {
	cfgRaw, ok := extra[pluginName].(map[string]interface{})
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

func main() {}

// This logger is replaced by the RegisterLogger method to load the one from KrakenD
var logger Logger = noopLogger{}

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] OTEL Handler Plugin Loaded", pluginName))
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
