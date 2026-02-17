#/bin/bash

# This add the CA to the server keystore (for kafka), 
# when making requests, we either need to add the 
# CA to the system valid CA's or if using java, adding 
# it to its keystore too
echo "              "
echo "--------------"
echo " Importing CA Cert to Keystore file:"
echo "   ${KEYSTORE_FILE}"
echo "--------------"
keytool \
    -keystore ${KEYSTORE_FILE} \
    -alias KrakenDCARoot \
    -import \
    -file cacert.pem

# Generate the key pair for the server
echo "              "
echo "--------------"
echo " Generating Secure Kafka Server Key: "
echo "--------------"
keytool \
    -keystore ${KEYSTORE_FILE} \
    -alias localhost \
    -validity ${VALIDITY} \
    -genkey \
    -keyalg RSA \
    -storetype pkcs12

# Generate the signing request for the generated key pair
echo "              "
echo "--------------"
echo " Generating Secure Kafka Server SIGNING REQUEST: "
echo "--------------"
keytool \
    -keystore ${KEYSTORE_FILE} \
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
    -keyout ${CLIENT_CERT_KEYPAIR} \
    -out ${CLIENT_CERT_SIGN_REQUEST} 
