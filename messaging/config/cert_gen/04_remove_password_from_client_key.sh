#/bin/bash

openssl rsa \
    -in ${CLIENT_CERT_KEYPAIR} \
    -out ${CLIENT_CERT_KEYPAIR_PASSWORDLESS}
