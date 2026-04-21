package main

var ModifierRegisterer = respregisterer("streaming-modifier")

type respregisterer string

func (r respregisterer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
),
) {
	f(string(r), r.response, false, true)
}

func (r respregisterer) response(
	extra map[string]interface{},
) func(interface{}) (interface{}, error) {
	return func(input interface{}) (interface{}, error) {
		respw, ok := input.(ResponseWrapper)
		if !ok {
			return nil, errUnkownResponseType
		}

		return handleResponseStream(respw), nil
	}
}
