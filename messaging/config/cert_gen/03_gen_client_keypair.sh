#!/bin/bash

openssl \
    req -newkey rsa:2048 \
    -keyout ${CLIENT_CERT_KEYPAIR} \
    -out ${CLIENT_CERT_SIGN_REQUEST} 
