up:
	docker compose -p redis_gov up -d 
.PHONY: up

cluster: 
	redis-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003
.PHONY: cluster

down:
	docker compose -p redis_gov down
.PHONY: down
