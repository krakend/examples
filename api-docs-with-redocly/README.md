## API Docs generation with Redocly
This example takes a dummy KrakenD configuration and generates its Redocly documentation, as if the process was in a CI/CD pipeline.

In the configuration file you'll see examples on how to load markdown documentation, include external files, or add examples of rquests and responses.

**To test this example you need to copy your `LICENSE` file inside the `krakend` folder.**

Then you can build this image with:
```
docker build -t test --progress=plain --no-cache .
```

And start the server with:
```
docker run --rm -p "8080:8080" test
```

You can then see the generated documentation under [http://localhost:8080/docs/](http://localhost:8080/docs/)

## Docs generation with SwaggerUI
You can do the same thing with Swagger UI instead of Redocly by using the alternative `Dockerfile`. It only changes this step:

```
docker build -t test --progress=plain -f Dockerfile_swagger .
```
