ARG KRAKEND_BUILDER
ARG KRAKEND_VERSION

FROM krakend/builder-ee:$KRAKEND_BUILDER AS builder

WORKDIR /app
COPY --chown=krakend:nogroup plugins /app

RUN cd quota-control-mw && make amd64

FROM krakend/krakend-ee:$KRAKEND_VERSION

WORKDIR /etc/krakend

COPY config/krakend .
COPY LICENSE LICENSE
COPY --from=builder --chown=krakend:nogroup /app/*.so /opt/krakend/plugins/

CMD [ "run", "-c", "/etc/krakend/krakend.json" ]