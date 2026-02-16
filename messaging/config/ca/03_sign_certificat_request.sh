#/bin/bash

if [ ! -f "ca_database_index.txt" ]; then
    touch ca_database_index.txt
    echo "0001" > ca_serial.txt
fi

openssl \
    ca -config openssl-ca.cnf \
    -policy signing_policy \
    -extensions signing_req \
    -out ${CLIENT_SIGNED_CERT} \
    -infiles ${CLIENT_CERT_SIGN_REQUEST} 

#subjectKeyIdentifier   = hash
#     -keystore ${KEYSTORE_FILE} \
#     -alias localhost \
#     -import \
#     -file $1
