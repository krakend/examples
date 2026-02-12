# KrakenD Version
PLUGINAME := quota-control-mw
CROSSCOMPILER := aarch64-linux-musl-gcc
# Uncomment if using the linux-generic builder
#CROSSCOMPILER := aarch64-linux-gnu-gcc

go.mod:
	go mod init quota-control-mw
	go mod tidy

amd64:
	GOARCH=amd64 \
	ARCH=amd64 \
	CGO_ENABLED=1 \
	go build -buildmode=plugin -o ../${PLUGINAME}-amd64.so .

arm64:
	CC=${CROSSCOMPILER} \
	GOHOSTARCH=amd64 \
	GOARCH=arm64 \
	ARCH=arm64 \
	CGO_ENABLED=1 \
	go build -buildmode=plugin -ldflags="-extld=${CROSSCOMPILER}" -o ../${PLUGINAME}-arm64.so .
