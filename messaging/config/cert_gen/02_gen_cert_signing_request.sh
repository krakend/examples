#/bin/bash

# keytool \
#     -keystore ${KEYSTORE_FILE} \
#     -certreq \
#     -alias localhost \
#     -keyalg RSA \
#     -storetype pkcs12 \
#     -file ${SERVER_CERT_SIGN_REQUEST} \
#     -ext SAN=DNS:${SERVER_FQDN},IP:${SERVER_IPADDRESS}
