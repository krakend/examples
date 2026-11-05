# Generate CA and certificates

Run `make setup` from the parent directory. It runs [`setup.sh`](./setup.sh)
inside the `apache/kafka` image, which ships `keytool` and `openssl`, so you
don't need Java installed. The script runs unattended and generates:

- A self-signed CA in `ca/` (`cacert.pem` and `cakey.pem`).
- The secure broker certificate, valid for `localhost` and `seckafkabroker`,
  stored with the CA in `certs/keystore.jks`.
- The client certificate used by KrakenD and `kaf` for mTLS, in
  `certs/client/`.

You can change the settings in [`env.sh`](./env.sh). The most relevant one is
the keystore password (`ksp4ssword`).

To regenerate everything, run `make cleanup` and then `make setup` again.
