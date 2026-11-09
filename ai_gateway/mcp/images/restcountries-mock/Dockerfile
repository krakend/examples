FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /restcountries-mock .

FROM scratch
COPY --from=builder /restcountries-mock /restcountries-mock
EXPOSE 8080
ENTRYPOINT ["/restcountries-mock"]
