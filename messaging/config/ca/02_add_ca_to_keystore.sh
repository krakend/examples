#/bin/bash

# This add the CA to the server keystore (for kafka), 
# when making requests, we either need to add the 
# CA to the system valid CA's or if using java, adding 
# it to its keystore too
keytool \
    -keystore ${KEYSTORE_FILE} \
    -alias KrakenDCARoot \
    -import \
    -file cacert.pem

