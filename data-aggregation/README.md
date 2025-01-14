###### KrakenD - Quick Video Tutorials

# Data aggregation with KrakenD

<div align="center">

| <a href="https://youtu.be/Otxcyy9f3bI"><img src="https://i.ytimg.com/vi/Otxcyy9f3bI/maxresdefault.jpg" width="300" heigth="300"></a><br>[Video in English](https://youtu.be/Otxcyy9f3bI) | <a href="https://youtu.be/KklqX8zqhAI"><img src="https://i.ytimg.com/vi/KklqX8zqhAI/maxresdefault.jpg" width="300" heigth="300"></a><br>[Video in Spanish](https://youtu.be/KklqX8zqhAI) |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|

</div>

In this repository, you will find the example code for the video tutorial on how to use KrakenD to aggregate data from multiple sources into a single endpoint. This is a powerful and versatile functionality that can help you present data from multiple sources in a cohesive way.

## Getting Started

To get started with the example code, you will need to have KrakenD installed on your system. You can follow the instructions in the [KrakenD documentation](hhttps://www.krakend.io/docs/overview/) to install it.

Once you have KrakenD installed, you can clone this repository to your local machine and navigate to the examples directory.

```bash
$ git clone https://github.com/krakend/examples.git
$ cd 5.data-aggregation
```

Inside the `5.data-aggregation` directory, you will find the configuration file `krakend.json` that you can use to run the examples.

To start the example, simply run the following command:

```
$ docker-compose up
```

This will start the KrakenD API gateway and you will be able to access the endpoint at http://localhost:8080/dashboard.

## Example Config Explanation

In the configuration file krakend.json, you can find the definition of the endpoint `/dashboard` that aggregates data from three different services concurrently.

```json
"endpoints": [
  {
    "endpoint": "/dashboard",
    "backend": [
      {
        "url_pattern": "/customers",
        "host": "http://my-service:8080"
      },
      {
        "url_pattern": "/sales",
        "host": "http://my-service:8080"
      },
      {
        "url_pattern": "/inventory",
        "host": "http://my-service:8080"
      }
    ]
  }
]
```

This is a simple example, but you can find more advanced configurations in the KrakenD documentation.

# Conclusion

We hope this example code and tutorial have been helpful in showing you how to use KrakenD for data aggregation. Don't hesitate to ask questions or provide feedback in the repository's issues section.
