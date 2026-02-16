# Generate CA and certificates

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
