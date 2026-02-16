#/bin/bash

export SERVER_CERT_SIGN_REQUEST=${PWD}/certs/server/localhost.csr
export SERVER_SIGNED_CERT=${PWD}/certs/server/localhost.signed.pem
export SERVER_FQDN=localhost
export SERVER_IPADDRESS=127.0.0.1

export CLIENT_CERT_KEYPAIR=${PWD}/certs/client/client.key
export CLIENT_CERT_SIGN_REQUEST=${PWD}/certs/client/client.csr
export CLIENT_SIGNED_CERT=${PWD}/certs/client/client.signed.pem
export CLIENT_CERT_KEYPAIR_PASSWORDLESS=${PWD}/certs/client/client.passwordless.key

export KEYSTORE_FILE=${PWD}/certs/keystore.jks
export VALIDITY=3650

