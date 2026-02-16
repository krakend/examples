#/bin/bash

# Create the "root" CA keypair / certificate
openssl \
    req -x509 \
    -config openssl-ca.cnf \
    -newkey rsa:4096 \
    -sha256 \
    -nodes \
    -out cacert.pem \
    -outform PEM
