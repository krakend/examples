# Generate CA and certificates

To generate the example you should use these passwords:

- for keystore: `ksp4ssword`
- for client certificate: `cl1entpass`

You can select other passwords, but then you will need to change the 
passwords in: 

- `certs/keystore.credentials.txt`: for the keystore
- `clients/producers/ka_fssl_client_cert_password`: for the client certificate


1. Edit the `env.sh` file, and check the environment vars 
    to know where the files will be placed.
    
2. Source the `env.sh` file
   
```bash
source ./env.sh
```

1. Enter the `ca` dir, to generate the Certificate Authority
    - `./01_gen_va.sh`
    - `./02_add_ca_to_keystore.sh`

2. Go  to the `cert_gens` dir, and create a certificate, 
   with its signing request.


    - `./03_sign_certificat_request.sh`

**Edit the `certs/keystore.credentials.txt`** and change the placeholder
password `my_password` for the password used for the keystore. This
file is used by the dockerized secure kafka instance to read the 
"secret" password for the keystore.
