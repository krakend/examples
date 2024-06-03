# Multiple POST Example with SOAP Backends in KrakenD

This repository contains a configuration file for [KrakenD](https://www.krakend.io/). The aim of this configuration is to illustrate how KrakenD can aggregate multiple POST calls into a single endpoint. This specific example also demonstrates the transformation from XML to JSON, given that the backend services are SOAP based.

## Configuration Overview

The KrakenD configuration in this repository is designed to interact with multiple SOAP backend services hosted at `http://webservices.oorsprong.org`. Specifically, the services are used to gather information about a country including the capital city, currency, and flag image.

The aggregation is handled by a single KrakenD endpoint: **`/test`**. This endpoint aggregates data from three backend SOAP services:

- **Capital City Service**: Retrieves the capital city of a country.
- **Country Currency Service**: Retrieves the country's currency name and symbol.
- **Country Flag Service**: Retrieves the country's flag image.

Each service receives a static POST request with a specific SOAP envelope as the body. This SOAP envelope is base64-encoded and is included in the configuration. Since the bodies are static, they are always requesting information about Australia (Country ISO Code "AU").

The JSON response from the `/test` endpoint is as follows:

```json
{
	"capital":"Canberra",
	"currency_name":"Australian Dollars",
	"currency_symbol":"AUD",
	"flag":"http://www.oorsprong.org/WebSamples.CountryInfo/Flags/Australia.jpg"
}
```

## KrakenD Enterprise and SOAP Integration

This configuration is limited in the sense that it uses static POST bodies for the SOAP requests. If you need to create dynamic POST bodies to send different SOAP requests based on the received REST parameters, consider using the [KrakenD Enterprise Edition](https://www.krakend.io/docs/enterprise/backends/soap/) which includes advanced SOAP integration capabilities.

## Running KrakenD with this Configuration

Assuming you have KrakenD installed, you can run it with this configuration using the following command:

```bash
krakend run -c krakend.json
```

Once KrakenD is running, you can retrieve the aggregated country information using the following CURL command:

```bash
curl http://localhost:8080/test
```

Or simply visit http://localhost:8080/test in your web browser.

## Conclusion

This repository illustrates the capabilities of KrakenD for aggregating multiple POST requests to SOAP backend services into a single endpoint, and converting XML responses to JSON. This is particularly useful when dealing with multiple SOAP services and trying to provide a unified and RESTful interface to your clients.
