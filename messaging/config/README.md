# Generate CA and certificates

To generate the example just run `bash ./setup.sh`, and fill the missing
fields for certificates (and accept signing and adding the certificates 
to the keystore).

You can edit the `env.sh` file to change settins if you want, however,
the only interesting variable to look at is the keystore password
(that is set to `ksp4ssword`).
    
After genrating the self signed CA you can start the docker compose environment.
