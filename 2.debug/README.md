###### KrakenD - Quick Video Tutorials

# Debugging KrakenD with the debug endpoint

<div align="center">

| <a href="https://youtu.be/oUZUlI8I-v8"><img src="https://i.ytimg.com/vi/oUZUlI8I-v8/maxresdefault.jpg" width="300" heigth="300"></a><br>[Video in English](https://youtu.be/oUZUlI8I-v8) | <a href="https://youtu.be/oFiYT3GPu_E"><img src="https://i.ytimg.com/vi/oFiYT3GPu_E/maxresdefault.jpg" width="300" heigth="300"></a><br>[Video in Spanish](https://youtu.be/oFiYT3GPu_E) |
|---|---|

</div>

## The `/__debug` endpoint

KrakenD provides a debugging service, under the endpoint `__debug`. This service can be used as a backend for any endpoint in your configuration. Once this is done, you will see that, when calling that endpoint, the KrakenD log will show detailed information, not about the request that you have sent to your gateway but about the request that your gateway is sending to your backend.

Keeping this in mind is important since KrakenD operates under a _zero-trust policy_ and [does not forward any query string or header without your explicit permission](https://www.krakend.io/docs/endpoints/parameter-forwarding/).

In this example, you'll see that we have a `/my-endpoint` endpoint, that allows forwarding of the parameter `content_id` as a querystring:

```json
    {
      "endpoint": "/my-endpoint",
      "backend": [
        {
          "host": ["http://localhost:8080"],
          "url_pattern": "/__debug/"
        }
      ],
      "input_query_strings": ["content_id"]
    }
```

To test it, just clone this repo and run KrakenD from this `2.debug` folder. You can run KrakenD using Docker, mapping the current folder to make the configuration file available for KrakenD inside the container, as follows:

```shell
$ docker run \
--rm -it -p "8080:8080" \
-v "$PWD:/etc/krakend" \
devopsfaith/krakend run -dc "krakend.json"
```

Once KrakenD is up & running, you can send a test request from other terminal tab:

```shell
$ curl -iG "http://localhost:8080/my-endpoint?content_id=1&unknown_param=2"
```

If you take a look at the output logs on the terminal tab where KrakenD is running, you'll see something like this:

```shell
KRAKEND DEBUG: [ENDPOINT: /__debug/*] Method: GET
KRAKEND DEBUG: [ENDPOINT: /__debug/*] URL: /__debug/?content_id=1
KRAKEND DEBUG: [ENDPOINT: /__debug/*] Query: map[content_id:[1]]
KRAKEND DEBUG: [ENDPOINT: /__debug/*] Params: [{param /}]
KRAKEND DEBUG: [ENDPOINT: /__debug/*] Headers: map[Accept-Encoding:[gzip] User-Agent:[KrakenD Version 2.0.4] X-Forwarded-For:[127.0.0.1] X-Forwarded-Host:[localhost:8080]]
KRAKEND DEBUG: [ENDPOINT: /__debug/*] Body: 
KRAKEND DEBUG: [GIN] 200 |     128.292µs |             ::1 | GET      "/__debug/?content_id=1"
KRAKEND DEBUG: [GIN] 200 |     7.17425ms |       127.0.0.1 | GET      "/my-endpoint?content_id=1&unknown_param=2"
```

Note that even if we sent an `unknown_param=2` in the original request, only `content_id` reached our backend, since it's the only param explicitly authorized in the configuration.

You will find detailed information in our documentation, as always: https://www.krakend.io/docs/endpoints/debug-endpoint/
