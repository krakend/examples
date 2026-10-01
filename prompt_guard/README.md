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

There is an endpoint that produces random results with the payload 
of a classifier in the "fake api" server. 

I can be tested directly with this request:

```bash
curl -i http://localhost:8088/scorer/
```

So at the service level configuration, there is prompt guard of type 
`classifier` named `random_scorer` that uses that server to 
decide if a payload should be served or not.

## Policy Prompt Guard

At the service level there is a definition of a guard named `max_size_policy`
that uses CEL expressions to block payloads that are too big or that 


## The metrics

Traces can be checked at [http://localhost:16686](http://localhost:16686).

And metrics at [http://localhost:4000](http://localhost:4000) with user: `krakend`
and password `krakend`. 

