setup:
	cd ./config && bash ./setup.sh
	docker compose up -d
.PHONY: setup

cleanup:
	docker compose down
	cd ./config && bash ./cleanup.sh
.PHONY: cleanup

up:
	docker compose up -d
.PHONY: up

down:
	docker compose down
.PHONY: down

run: 
	./krakend run -c ./config/krakend.json
.PHONY: run

