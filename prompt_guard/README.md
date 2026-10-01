# KrakenD Prompt Guard Example

This is an example about how to use the [AI Gateway Prompt Guard](https://www.krakend.io/docs/enterprise/ai-gateway/prompt-guard/)
component. 

Despite the fact this component is targeted at AI body payloads,
it can used on any `POST` / `PUT` endpoint.

## Default Prompt Guards

In the `config/krakend/krakend.json` file you can find the 
`/defult_guards/` endpoint configuration example, with a list of empty guards, 
that means only defaults will be applied.

In the `curl.sh` there are 3 requests made that shows how no poisoned data
can get the result back, while some other prompts / bodies will be 
blocked by the default set of guards.

## Regex Prompt Guard 

There are 2 endpoints, for the custom regex guards, to showcase that 
the filtering can be put at the endpoint level or at the backend level:

- `/personal/at_backend`
- `/personal/at_endpoint`

Both umake use of the `email` and `id_card` regex guards defined
at the `"ai/prompt_guard"` service level

## External classifier Prompt Guard

## Policy Prompt Guard

## The metrics
