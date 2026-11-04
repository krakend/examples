FROM golang:1.25.7-alpine3.23

WORKDIR /app
COPY openai-mock .
RUN go build -o openai-mock .

EXPOSE 8090
CMD ["/app/openai-mock"]