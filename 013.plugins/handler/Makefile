# KrakenD Version
VERSION := latest
PLUGINAME := plugins/handler

# Build KrakenD plugin using the KrakenD docker builder (for running inside docker)
amd64:
	docker run --rm -it -e CGO_ENABLED=1 -e GOARCH=amd64 -e ARCH=amd64 -v "${PWD}:/app" -w /app krakend/builder:${VERSION} sh -c "make -e plugin"

arm64:
	docker run --rm -it -e CGO_ENABLED=1 -e CC=aarch64-linux-musl-gcc -e GOHOSTARCH=amd64 -e GOARCH=arm64 -e ARCH=arm64 -e EXTRA_LDFLAGS='-extld=aarch64-linux-musl-gcc' -v "${PWD}:/app" -w /app krakend/builder:${VERSION} sh -c "make -e plugin"

# Build KrakenD plugin using the KrakenD docker builder for non-docker plugins
linux_amd64:
	docker run --rm -it -e CGO_ENABLED=1 -e GOARCH=amd64 -e ARCH=amd64 -v "${PWD}:/app" -w /app krakend/builder:${VERSION}-linux-generic sh -c "make -e plugin"

linux_arm64:
	docker run --rm -it -e CGO_ENABLED=1 -e CC=aarch64-linux-gnu-gcc -e ARCH=arm64 -e GOARCH=arm64 -e GOHOSTARCH=amd64 -e EXTRA_LDFLAGS='-extld=aarch64-linux-gnu-gcc' -v "${PWD}:/app" -w /app krakend/builder:${VERSION}-linux-generic sh -c "make -e plugin"

plugin:
	@go build -buildmode=plugin -ldflags="${EXTRA_LDFLAGS}" -o ${PLUGINAME}.so .
