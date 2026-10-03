SHELL := /bin/bash

# Cert Authority from Skynet
ca:
	openssl req -newkey rsa:2048 -nodes -x509 -days 3650 -out ca.crt -keyout ca.key \
		-subj "/C=ES/O=SkyNet/CN=localhost"
.PHONY: ca

keycloak_cert:
	mkdir -p server
	openssl genrsa -out server/keycloak.key 2048
	openssl req -new \
		-key server/keycloak.key \
		-subj "/C=ES/O=KeyCloak_Cert/CN=keycloak" \
		-addext "subjectAltName = DNS:keycloak" \
		-out server/keycloak.csr 
	openssl x509 -req -in server/keycloak.csr \
		-CA ca.crt -CAkey ca.key -CAcreateserial \
		-extfile <(printf "subjectAltName=DNS:keycloak,DNS:127.0.0.1") \
		-out server/keycloak.crt -days 3650 -sha256
.PHONY: keycloak_cert

krakend_cert:
	mkdir -p server
	openssl genrsa -out server/krakend_ee.key 2048
	openssl req -new \
		-key server/krakend_ee.key \
		-subj "/C=ES/O=Krakend_EE_Cert/CN=krakend_ee" \
		-addext "subjectAltName=DNS:krakend_ee" \
		-out server/krakend_ee.csr 
	openssl x509 -req -in server/krakend_ee.csr \
		-CA ca.crt -CAkey ca.key -CAcreateserial \
		-extfile <(printf "subjectAltName=DNS:krakend_ee,DNS:127.0.0.1") \
		-out server/krakend_ee.crt -days 3650 -sha256
.PHONY: krakend_cert

all: ca keycloak_cert krakend_cert
