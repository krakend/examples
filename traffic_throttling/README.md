# KrakenD - Return error details

Test environment to test some traffic throttling features of KrakenD.

You can run the environment with `docker-compose up`. To follow KrakenD logs on terminal you can run `docker-compose logs -f krakend`

### Root Endpoint

http://localhost:8080/ 👈 Returns the default answer from internal backend (exposed at http://localhost:8001/).

**Example call**

Here we're sending a single call to root endpoint. 

```shell
$ curl -iG http://localhost:8080/
               
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Krakend: Version 2.0.5
X-Krakend-Completed: true
Date: Thu, 23 Jun 2022 10:59:11 GMT
Content-Length: 25

{"message":"Hello world"}
```
### Endpoint with rate limit

http://localhost:8080/rate-limit 👈 This endpoint points to the same internal backend and url_pattern that the root endpoint, but applying a rate limit with a max rate of `1`. You can read all available options at https://www.krakend.io/docs/endpoints/rate-limit/ Rate limit can be applied also at backend (proxy) level: https://www.krakend.io/docs/backends/rate-limit/

**Example**
Here we're sending to sequential calls to the rate limited endpoint.
```shell
$ curl -iG http://localhost:8080/rate-limit && echo "\n\n------\nSending a new request\n------\n" && curl -iG http://localhost:8080/rate-limit

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Krakend: Version 2.0.5
X-Krakend-Completed: true
Date: Thu, 23 Jun 2022 11:03:32 GMT
Content-Length: 25

{"message":"Hello world"}

------
Sending a new request
------

HTTP/1.1 503 Service Unavailable
Date: Thu, 23 Jun 2022 11:03:32 GMT
Content-Length: 0
```

The first call is correctly served. The second call has been blocked, and a 503 Service Unavailable status code has been returned.

### Endpoint with circuit breaker

http://localhost:8080/circuit-breaker 👈 This endpoint points has two backends defined that will be called concurrently and grouped in "backedn1" and "backend2" groups, respectively. We'll be defined a behaviour on the service called that will provoke a random failure (500 status code) of the service on 25% of requests.

We've configured the circuit breaker to prevent sending further requests to the failing backend when KrakenD detects 1 error in a 5 seconds interval, keeping circuit opened for 2 seconds before trying sending requests again to the failing backend.

**Example 1**

This example shows a call where both backends answer correctly.
```shell
$ curl -iG http://localhost:8080/circuit-breaker                               
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Krakend: Version 2.0.5
X-Krakend-Completed: true
Date: Thu, 23 Jun 2022 11:10:10 GMT
Content-Length: 75

{"backend1":{"message":"Hello world"},"backend2":{"message":"Hello world"}}
```
Important to notice here the two answers grouped in "backend1" and "backend2", and the "X-KrakenD-Completed" header, that indicates that the gateway was able to collect all needed data.

**Example 2**

This example shows a call where one backend fails on first request. On a second request, the call will not be able to reach the backend, since it will find an "opened circuit".

```shell
$ curl -iG http://localhost:8080/circuit-breaker
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Krakend: Version 2.0.5
X-Krakend-Completed: false
Date: Thu, 23 Jun 2022 11:13:58 GMT
Content-Length: 38

{"backend1":{"message":"Hello world"}}
```

Note that the "backend2" group has not been returned (second backend failed) and the "X-KrakenD-Completed" header was set to false, indicating that KrakenD wasn't able to collect all needed data from defined backends.

On KrakenD logs, we will see a message noticing this problem:

```shell
krakend_1   | 2022/06/23 11:16:47 KRAKEND ERROR: [ENDPOINT: /circuit-breaker] Error #0: invalid status code
krakend_1   | [GIN] 2022/06/23 - 11:16:47 | 200 |    6.473959ms |      172.18.0.1 | GET      "/circuit-breaker"
```

If we try to send more requests before the configured retry interval, we will see that those requests will be blocked by the circuit breaker:

```shell
krakend_1   | 2022/06/23 11:16:48 KRAKEND WARNING: [CB] Circuit breaker named 'test-circuit-breaker' went from 'closed' to 'open'
krakend_1   | 2022/06/23 11:16:48 KRAKEND ERROR: [ENDPOINT: /circuit-breaker] Error #0: invalid status code
krakend_1   | [GIN] 2022/06/23 - 11:16:48 | 200 |     12.7625ms |      172.18.0.1 | GET      "/circuit-breaker"
krakend_1   | 2022/06/23 11:16:49 KRAKEND ERROR: [ENDPOINT: /circuit-breaker] Error #0: circuit breaker is open
krakend_1   | [GIN] 2022/06/23 - 11:16:49 | 200 |    3.794042ms |      172.18.0.1 | GET      "/circuit-breaker"
krakend_1   | 2022/06/23 11:16:50 KRAKEND ERROR: [ENDPOINT: /circuit-breaker] Error #0: circuit breaker is open
```

If the backend recovers and start answering without errors, the circuit breaked will be closed again, allowing the failing backedn to receive requests again:

```shell
krakend_1   | 2022/06/23 11:16:50 KRAKEND WARNING: [CB] Circuit breaker named 'test-circuit-breaker' went from 'open' to 'half-open'
krakend_1   | 2022/06/23 11:16:50 KRAKEND WARNING: [CB] Circuit breaker named 'test-circuit-breaker' went from 'half-open' to 'closed'
krakend_1   | [GIN] 2022/06/23 - 11:16:50 | 200 |   10.622916ms |      172.18.0.1 | GET      "/circuit-breaker"
```
