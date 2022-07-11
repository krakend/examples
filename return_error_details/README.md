# KrakenD - Return error details

Test environment to validate `X-KrakenD-Completed` headers and `return_error_details` KrakenD config parameter.

You can run the environment with `docker-compose up`

### Complete answer ✅

http://localhost:8080/ok 👈 Gives you a complete answer coming from two backends with a valid JSON response and 200 status

**Body**
```json
{
  "a": {
    "message": "Hello world"
  },
  "b": {
    "message": "Hello world"
  }
}
```

**Headers**
```ini
Content-Length: 61
Content-Type: application/json; charset=utf-8
Date: Thu, 10 Mar 2022 21:37:15 GMT
X-Krakend: Version 2.0.0
X-Krakend-Completed: true
```
### Incomplete answer ⚠️

http://localhost:8080/ko 👈 Gives you an incomplete answer coming from two backends: one with a valid JSON and 200 status, and another returning a 500 status code.

**Body**
```json
{
  "a": {
    "message": "Hello world"
  },
  "error_backend_GetSubscriptionBalanceResponse": {
    "http_status_code": 500
  }
}
```

**Headers**
```ini
Content-Length: 103
Content-Type: application/json; charset=utf-8
Date: Thu, 10 Mar 2022 21:38:10 GMT
X-Krakend: Version 2.0.0
X-Krakend-Completed: false
```
