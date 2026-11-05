# The certificates are generated inside the apache/kafka image, which ships
# keytool and openssl, so you don't need Java installed locally.
setup:
	docker run --rm --user "$$(id -u):$$(id -g)" -v "$(CURDIR)/config:/config" -w /config --entrypoint bash apache/kafka:latest ./setup.sh
.PHONY: setup

cleanup:
	docker compose down
	bash ./config/cleanup.sh
.PHONY: cleanup

up:
	sh ./startup.sh
.PHONY: up

down:
	docker compose down
.PHONY: down
