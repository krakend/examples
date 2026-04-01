package main

import (
	"context"
)

var MiddlewareRegisterer = mwregisterer("streaming-middleware")

type mwregisterer string

func (r mwregisterer) RegisterMiddlewares(f func(
	name string,
	middlewareFactory func(map[string]interface{}, func(context.Context, interface{}) (interface{}, error)) func(context.Context, interface{}) (interface{}, error),
),
) {
	f(string(r), r.middlewareFactory)
}

func (r mwregisterer) middlewareFactory(cfg map[string]interface{}, next func(context.Context, interface{}) (interface{}, error)) func(context.Context, interface{}) (interface{}, error) {
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

		return handleResponseStream(respw), nil
	}
}
