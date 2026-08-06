#!/bin/bash

mkdir -p ./certs/server
mkdir -p ./certs/client

. env.sh

echo ${KEYSTORE_PASS} > ./certs/keystore.credentials.txt

cd ca
echo "              "
echo "--------------"
echo " Creating CA Cert:"
echo "--------------"
openssl \
    req -x509 \
    -config openssl-ca.cnf \
    -newkey rsa:4096 \
    -sha256 \
    -nodes \
    -out cacert.pem \
    -outform PEM

# . ./01_gen_ca.sh
# . ./02_add_ca_to_keystore.sh
echo "              "
echo "--------------"
echo " Importing CA Cert to Keystore file:"
echo "   ${KEYSTORE_FILE}"
echo "--------------"
keytool \
    -keystore ${KEYSTORE_FILE} \
    -storepass ${KEYSTORE_PASS} \
    -alias KrakenDCARoot \
    -import \
    -file cacert.pem

echo "              "
echo "--------------"
echo " Generating Secure Kafka Server Key: "
echo "--------------"
keytool \
    -keystore ${KEYSTORE_FILE} \
    -storepass ${KEYSTORE_PASS} \
    -alias localhost \
    -validity ${VALIDITY} \
    -genkey \
    -keyalg RSA \
    -storetype pkcs12

echo "              "
echo "--------------"
echo " Generating Secure Kafka Server SIGNING REQUEST: "
echo "--------------"
keytool \
    -keystore ${KEYSTORE_FILE} \
    -storepass ${KEYSTORE_PASS} \
    -certreq \
    -alias localhost \
    -keyalg RSA \
    -storetype pkcs12 \
    -file ${SERVER_CERT_SIGN_REQUEST} \
    -ext SAN=DNS:${SERVER_FQDN},IP:${SERVER_IPADDRESS}

echo "              "
echo "--------------"
echo " Generating Client Key Pair and SIGNING REQUEST: "
echo "--------------"
openssl \
    req -newkey rsa:2048 \
    -subj '/CN=localhost/OU=TEST/O=KrakenD/L=Girona/ST=Girona/C=ES' \
    -config ./san.cnf \
    -nodes \
    -keyout ${CLIENT_CERT_KEYPAIR} \
    -out ${CLIENT_CERT_SIGN_REQUEST} 


echo "              "
echo "--------------"
echo " Checking CA DB exists: "
echo "--------------"
if [ ! -f "ca_database_index.txt" ]; then
    touch ca_database_index.txt
    echo "0001" > ca_serial.txt
fi

echo "              "
echo "--------------"
echo " Signing Server CERT: "
echo "--------------"
openssl \
    ca -config openssl-ca.cnf \
    -policy signing_policy \
    -extensions signing_req \
    -out ${SERVER_SIGNED_CERT} \
    -infiles ${SERVER_CERT_SIGN_REQUEST} 

echo "              "
echo "--------------"
echo " Import Server CERT into keystore:"
echo "--------------"
keytool \
    -importcert \
    -keystore ${KEYSTORE_FILE} \
    -storepass ${KEYSTORE_PASS} \
    -alias localhost \
    -file ${SERVER_SIGNED_CERT}

echo "              "
echo "--------------"
echo " Signing Client CERT: "
echo "--------------"
openssl \
    ca -config openssl-ca.cnf \
    -policy signing_policy \
    -extensions signing_req \
    -out ${CLIENT_SIGNED_CERT} \
    -infiles ${CLIENT_CERT_SIGN_REQUEST} 

cd ..
