#/bin/bash

# Create the "root" CA keypair / certificate
echo "              "
echo "--------------"
echo " Creating CA certificate:"
echo "--------------"
openssl \
    req -x509 \
    -config openssl-ca.cnf \
    -newkey rsa:4096 \
    -sha256 \
    -nodes \
    -out cacert.pem \
    -outform PEM
