#!/bin/bash

# openssl \
#     req -newkey rsa:2048 \
#     -keyout ${CLIENT_CERT_KEYPAIR} \
#     -out ${CLIENT_CERT_SIGN_REQUEST} 
# 
openssl \
    req -newkey rsa:2048 \
    -subj '/CN=localhost/OU=TEST/O=KrakenD/L=Girona/ST=Girona/C=ES' \
    -keyout ${CLIENT_CERT_KEYPAIR} \
    -out ${CLIENT_CERT_SIGN_REQUEST} \
    -nodes \
    -config ./san.cnf
