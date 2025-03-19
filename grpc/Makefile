docker_up:
	cd ./images/grpc && make bin_proto_multiple_files
	docker compose up -d 
.PHONY: docker_up

docker_stop:
	docker compose stop
.PHONY: docker_stop

docker_down:
	docker compose down
.PHONY:	docker_down


