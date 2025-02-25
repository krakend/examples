FROM golang:1.20-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN apk add git
RUN go get -d -v ./...
RUN go build ./cmd/server

FROM scratch
COPY --from=builder /app/server /server
COPY --from=builder /app/certs /certs
EXPOSE 4243
ENTRYPOINT ["/server"]
