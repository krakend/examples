ARG KRAKEND_BUILDER
ARG KRAKEND_VERSION

FROM --platform=linux/amd64 krakend/builder-ee:latest AS builder

WORKDIR /app
COPY --chown=krakend:nogroup plugins /app

RUN cd streaming-modifier && make arm64

FROM krakend/krakend-ee:latest

WORKDIR /etc/krakend

COPY config/krakend .
COPY --from=builder --chown=krakend:nogroup /app/*.so /opt/krakend/plugins/

CMD [ "run", "-c", "/etc/krakend/krakend.json" ]