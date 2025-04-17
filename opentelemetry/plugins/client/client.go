// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	otelclient "github.com/krakend/krakend-otel/http/client"

	"github.com/krakend/krakend-otel/state"
	"go.opentelemetry.io/otel/attribute"
)

var (
	pluginName            = "otel-client"
	ClientRegisterer      = registerer(pluginName)
	errRegistererNotFound = fmt.Errorf("%s plugin disabled: config not found", pluginName)
)

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r registerer) registerClients(_ context.Context, extra map[string]interface{}) (http.Handler, error) {
	// check the cfg. If the modifier requires some configuration,
	// it should be under the name of the plugin. E.g.:
	/*
	   "extra_config":{
	       "plugin/http-client":{
	           "name":"my-client-plugin",
	           "my-client-plugin":{
	               "option": "/some-path"
	           }
	       }
	   }
	*/
	config, err := parseConfig(extra)
	if err != nil {
		return nil, err
	}

	// we need to get the global otel state in order to set it in the transport options:
	otelConf := state.GlobalConfig()
	otelState := otelConf.OTEL() // get the global configured settings
	otelOpts := &otelclient.TransportOptions{
		MetricsOpts: otelclient.TransportMetricsOptions{
			RoundTrip:          true,
			ReadPayload:        true,
			DetailedConnection: true,
			FixedAttributes: []attribute.KeyValue{
				attribute.Key("attr_plugin_otel_client").String("something else"),
			},
		},
		TracesOpts: otelclient.TransportTracesOptions{
			RoundTrip:          true,
			ReadPayload:        true,
			DetailedConnection: true,
			FixedAttributes: []attribute.KeyValue{
				attribute.Key("attr_plugin_otel_client").String("trace_attr"),
			},
			ReportHeaders: true,
		},
		OTELInstance: otelState,
	}

	logger.Debug(fmt.Sprintf("The plugin is now running %s", config.Opt))

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		c := otelclient.InstrumentedHTTPClient(http.DefaultClient, otelOpts, "otelplugin")
		// use c for any requests you want to perform
		resp, err := c.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy headers, status codes, and body from the backend to the response writer
		for k, hs := range resp.Header {
			for _, h := range hs {
				w.Header().Add(k, h)
			}
		}
		w.WriteHeader(resp.StatusCode)
		if resp.Body == nil {
			return
		}
		io.Copy(w, resp.Body)
		resp.Body.Close()

	}), nil
}

type config struct {
	Opt string `json:"option"`
}

// parseConfig parses the configuration marshaling and unmarshaling into a struct.
// you can also manually check for fields in the extra conffig map
func parseConfig(extra map[string]interface{}) (*config, error) {
	if name, ok := extra["name"].(string); !ok || name != pluginName {
		return nil, errRegistererNotFound
	}

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

var logger Logger = nil

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Example client plugin loaded", pluginName))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}
