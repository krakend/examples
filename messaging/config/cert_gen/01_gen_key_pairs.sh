#/bin/bash

# Generate the key pair for the server
keytool \
    -keystore ${KEYSTORE_FILE} \
    -alias localhost \
    -validity ${VALIDITY} \
    -genkey \
    -keyalg RSA \
    -storetype pkcs12

# Generate the signing request for the generated key pair
keytool \
    -keystore ${KEYSTORE_FILE} \
    -certreq \
    -alias localhost \
    -keyalg RSA \
    -storetype pkcs12 \
    -file ${SERVER_CERT_SIGN_REQUEST} \
    -ext SAN=DNS:${SERVER_FQDN},IP:${SERVER_IPADDRESS}
