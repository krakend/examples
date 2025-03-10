if [ -z $DD_API_KEY ]; then
    echo -e "\nNo DD_API_KEY found. Set it with:\n\n    export DD_API_KEY="YOUR_API_KEY"\n"
    exit -1
fi


if [ -z $DD_SITE ]; then
    export DD_SITE=datadoghq.com
    # use datadoghq.eu for european accounts
fi

# docker run --name dd-agent --rm \
#     -e DD_SITE="datadoghq.com" \
#     -e DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_GRPC_ENDPOINT="0.0.0.0:4317" \
#     -e DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_HTTP_ENDPOINT="0.0.0.0:4318" \
#     -e DD_API_KEY=${DD_API_KEY} \
#     -e DD_APM_ENABLED=true \
#     -p 14317:4317 \
#     -p 14318:4318 \
#     -v /var/run/docker.sock:/var/run/docker.sock:ro \
#     -v /proc/:/host/proc/:ro \
#     -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
#     -v /var/lib/docker/containers:/var/lib/docker/containers:ro \
#     gcr.io/datadoghq/agent:7

# -e DD_HOSTNAME="localhost" \
# -e DD_CONTAINER_INCLUDE="" \
docker run --name dd-agent --rm \
    -e DD_SITE=$DD_SITE \
    -e DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_GRPC_ENDPOINT="0.0.0.0:4317" \
    -e DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_HTTP_ENDPOINT="0.0.0.0:4318" \
    -e DD_API_KEY=$DD_API_KEY\
    -e DD_APM_ENABLED=true \
    -e DD_APM_NON_LOCAL_TRAFFIC=true \
    -p 14317:4317 \
    -p 14318:4318 \
    -v /var/run/docker.sock:/var/run/docker.sock:ro \
    -v /proc/:/host/proc/:ro \
    -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
    -v /var/lib/docker/containers:/var/lib/docker/containers:ro \
  gcr.io/datadoghq/agent:7


# docker run --name dd-agent --rm \
#     -e DD_SITE="datadoghq.com" \
#     -e DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_GRPC_ENDPOINT="0.0.0.0:4317" \
#     -e DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_HTTP_ENDPOINT="0.0.0.0:4318" \
#     -e DD_API_KEY=$DD_API_KEY \
#     -e DD_APM_ENABLED=true \
#     -p 14317:4317 \
#     -p 14318:4318 \
#     -v /var/run/docker.sock:/var/run/docker.sock:ro \
#     -v /proc/:/host/proc/:ro \
#     -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
#     -v /var/lib/docker/containers:/var/lib/docker/containers:ro \
#     gcr.io/datadoghq/agent:7
