#!/bin/bash

echo "Prompt Guard at Endpoint"
echo "should be served"
echo ""
curl \
    -X POST \
    -d '{"fo": "bar"}' \
    localhost:8080/block_at_endpoint
echo ""
echo "---"

echo "should be blocked"
echo ""
curl \
    -X POST \
    -d '{"foo": "bar"}' \
    localhost:8080/block_at_endpoint
echo ""
echo "---"


echo ""
echo "Prompt Guard at Backend"
echo "should be served"
echo ""
curl \
    -X POST \
    -d '{"fo": "bar"}' \
    localhost:8080/block_at_backend
echo ""
echo "---"

echo "should be blocked"
echo ""
curl \
    -X POST \
    -d '{"foo": "bar"}' \
    localhost:8080/block_at_backend
echo ""
echo "---"
