#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
CATALOG_DIR=$SCRIPT_DIR/../../images/grpc
CATALOG_FILE=$CATALOG_DIR/fullcatalog.pb

PAYLOAD_FILE=flight_payload.json
GRPC_METHOD=flight_finder.Flights/FindFlight

CURDIR=$(pwd)
cd $SCRIPT_DIR

if [ ! -f ${CATALOG_FILE} ]; then
    cd $CATALOG_DIR
    make bin_proto_single_file
    cd $SCRIPT_DIR
fi

if command -v jq 2>&1 > /dev/null
then
    echo -e "grpcurl -d @ -plaintext \n     -protoset ${CATALOG_FILE} \n     localhost:8080 ${GRPC_METHOD} < ${PAYLOAD_FILE} | jq"
    grpcurl -d @ -plaintext \
        -protoset ${CATALOG_FILE} \
        localhost:8080 \
        ${GRPC_METHOD} < ${PAYLOAD_FILE} | jq
else 
    grpcurl -d @ -plaintext \
        -protoset ${CATALOG_FILE} \
        localhost:8080 \
        ${GRPC_METHOD} < ${PAYLOAD_FILE}
fi
