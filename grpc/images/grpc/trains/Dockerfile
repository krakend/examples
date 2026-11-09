FROM golang:1.25-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN apk add git openssl
RUN go get -d -v ./...
RUN go build ./cmd/server
# Self-signed CA and server certificate for the TLS connection
RUN mkdir -p certs && \
    openssl req -x509 -newkey rsa:2048 -nodes -days 3650 \
        -subj "/CN=KrakenD Example Trains CA" \
        -keyout certs/ca.key -out certs/ca.crt && \
    openssl req -newkey rsa:2048 -nodes \
        -subj "/CN=grpc_trains" \
        -keyout certs/server.key -out certs/server.csr && \
    printf "subjectAltName=DNS:grpc_trains,DNS:localhost,IP:127.0.0.1\n" > certs/ext.cnf && \
    openssl x509 -req -in certs/server.csr -days 3650 -sha256 \
        -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial \
        -extfile certs/ext.cnf -out certs/server.cert

FROM scratch
COPY --from=builder /app/server /server
COPY --from=builder /app/certs /certs
EXPOSE 4243
ENTRYPOINT ["/server"]
