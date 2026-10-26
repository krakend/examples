#!/bin/bash

echo "check that we can get compressed response:"
curl -i --compressed localhost:8080/data
echo " "

echo "check that we can get non compressed response:"
curl -i --compressed localhost:8080/data
echo " "

echo "test that websockets work as expected (echo service):" 
websocat ws://localhost:8080/ws


