FROM krakend/krakend-ee AS builder

# This is an Enterprise feature, make sure to have the LICENSE in the krakend dir
COPY krakend /etc/krakend

RUN FC_DEBUG=true krakend check -d -t -c "krakend.json"
RUN krakend openapi export -c /tmp/krakend.json -o /tmp/openapi.json

# Redocly generation
FROM redocly/cli as redocly
COPY --from=builder /tmp/openapi.json /spec/swagger.json
# Outputs under /spec/redoc-static.html
RUN /usr/local/bin/redocly build-docs swagger.json

FROM krakend/krakend-ee AS production
COPY krakend/LICENSE /etc/krakend/LICENSE
# Copy the final configuration and the documentation
# which will be served from KrakenD itself under /docs/
COPY --from=builder /tmp/krakend.json /etc/krakend/krakend.json
COPY --from=redocly /spec/redoc-static.html /etc/krakend/docs/index.html
