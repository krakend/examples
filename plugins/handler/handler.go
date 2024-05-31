// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	// pluginName is the plugin name
	pluginName = "my-handler-plugin"
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
		return h, errRegistererNotFound
	}

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(w, req)
		logger.Debug("config value option: %s", config.Opt)
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
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Example handler plugin loaded", pluginName))
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
