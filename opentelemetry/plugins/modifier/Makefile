# KrakenD Version
VERSION := latest
PLUGINAME := plugins/modifier

# Build KrakenD plugin using the KrakenD docker builder (for running inside docker)
amd64:
	mkdir -p ./plugins
	docker run --rm -it \
		-e CGO_ENABLED=1 \
		-e GOARCH=amd64 \
		-e ARCH=amd64 \
		-e FOR_USER=$$(id -u) \
		-v "${PWD}:/app" \
		-w /app \
		krakend/builder:${VERSION} sh -c "make -e plugin"
	mkdir -p ../../config/krakend/plugins
	mv ./${PLUGINAME}.so ../../config/krakend/plugins

arm64:
	mkdir -p ./plugins
	docker run --rm -it \
		-e CGO_ENABLED=1 \
		-e CC=aarch64-linux-musl-gcc \
		-e GOHOSTARCH=amd64 \
		-e GOARCH=arm64 \
		-e ARCH=arm64 \
		-e FOR_USER=$$(id -u) \
		-e EXTRA_LDFLAGS='-extld=aarch64-linux-musl-gcc' \
		-v "${PWD}:/app" \
		-w /app \
		krakend/builder:${VERSION} sh -c "make -e plugin"
	mkdir -p ../../config/krakend/plugins
	mv ./plugin/${PLUGINANME} ../../config/krakend/plugins

plugin:
	@go get ./...
	@go build -buildmode=plugin \
		-ldflags="${EXTRA_LDFLAGS}" \
		-o ${PLUGINAME}.so . \
		&& chown ${FOR_USER}:${FOR_USER} ${PLUGINAME}.so
