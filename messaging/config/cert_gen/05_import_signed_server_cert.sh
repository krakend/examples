#!/bin/bash

keytool 
    -importcert \
    -keystore ${KEYSTORE_FILE} \
    -alias localhost \
    -file ${SERVER_SIGNED_CERT}
