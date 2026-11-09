ARG KRAKEND_BUILDER=3.0.3
ARG KRAKEND_VERSION=3.0.3

FROM --platform=linux/amd64 krakend/builder-ee:$KRAKEND_BUILDER AS builder

WORKDIR /app
COPY --chown=krakend:nogroup plugins /app

RUN cd streaming-modifier && make amd64

FROM krakend/krakend-ee:$KRAKEND_VERSION

WORKDIR /etc/krakend

COPY config/krakend .
COPY --from=builder --chown=krakend:nogroup /app/*.so /opt/krakend/plugins/

CMD [ "run", "-c", "/etc/krakend/krakend.json" ]