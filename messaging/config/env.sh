#/bin/bash

mkdir -p ./certs/server
mkdir -p ./certs/client

export SERVER_CERT_SIGN_REQUEST=${PWD}/certs/server/localhost.csr
export SERVER_FQDN=localhost
export SERVER_IPADDRESS=127.0.0.1

export CLIENT_CERT_KEYPAIR=${PWD}/certs/client/client.key
export CLIENT_CERT_SIGN_REQUEST=${PWD}/certs/client/client.csr
export CLIENT_SIGNED_CERT=${PWD}/certs/client/client_cert.pem

export KEYSTORE_FILE=${PWD}/certs/keystore.jks
export VALIDITY=3650

echo "              "
echo "--------------"
echo "Creating CA:"
echo "--------------"

cd ca
. ./01_gen_ca.sh
. ./02_add_ca_to_keystore.sh
cd ..

echo "              "
echo "--------------"
echo "Creating Server Certificate"
echo "--------------"

cd cert_gen
. ./01_gen_key_pairs.sh
. ./03_gen_client_keypair.sh
cd ..

echo "              "
echo "--------------"
echo "Creating Client Certificate"
echo "--------------"
cd ca
. ./03_sign_certificat_request.sh $CLIENT_CERT_SIGN_REQUEST
cd ..
