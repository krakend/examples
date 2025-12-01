// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/krakend/krakend-otel/state"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
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
	// The config variable contains all the keys you have defined in the configuration
	// if the key doesn't exists or is not a map the plugin returns an error and the default handler
	config, err := parseConfig(extra)
	if err != nil {
		logger.Warning(fmt.Sprintf("[PLUGIN: %s] cannot parse config %s -> %#v",
			pluginName, err.Error(), extra))
		return h, errRegistererNotFound
	}

	// Get Access to the Global Open Telemetry config:
	otelConf := state.GlobalConfig()
	otelState := otelConf.OTEL() // get the global configured state
	tracer := otelState.Tracer()
	meter := otelState.Meter()

	instrumentName := "handler_plugin.default.duration"
	if config.Opt != "" {
		instrumentName = "handler_plugin." + config.Opt + ".duration"
	}
	fakeDurationInstrument, err := meter.Float64Histogram(instrumentName,
		metric.WithExplicitBucketBoundaries(0.010, 0.050, 0.100, 0.150))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rCtx := req.Context()
		span := trace.SpanFromContext(rCtx)

		// set attributes to current span
		myValue := "my_plugin_value"
		if config.Opt != "" {
			myValue = config.Opt
		}
		logger.Warning(fmt.Sprintf("[PLUGIN: %s] setting attr 'pluginvalue' = %s",
			pluginName, myValue))
		span.SetAttributes(attribute.Key("pluginvalue").String(myValue))

		// create a new span
		spanName := fmt.Sprintf("plugin-%s", pluginName)
		logger.Info(fmt.Sprintf("[PLUGIN: %s] starting new span: %s", pluginName, spanName))
		ctx, newSpan := tracer.Start(rCtx, spanName)
		req = req.WithContext(ctx)

		// fake doing something:
		rnd := rand.New(rand.NewSource(time.Now().UnixMicro()))
		// rndVal in seconds
		rndVal := float64(5+(rnd.Int63()%150)) / 1000.0
		fakeDurationInstrument.Record(ctx, rndVal)

		h.ServeHTTP(w, req)

		// end new span
		newSpan.End()
		logger.Info(fmt.Sprintf("[PLUGIN: %s] ending new span: %s", pluginName, spanName))

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
