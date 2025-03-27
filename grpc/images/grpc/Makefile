get_known_types:
	rm -rf ./contracts/google
	rm -rf ./tmp
	mkdir -p ./tmp && \
	 	cd ./tmp && \
	 	git clone --depth=1 https://github.com/protocolbuffers/protobuf.git
	mv ./tmp/protobuf/src/google ./contracts
	rm -rf ./contracts/google/protobuf/compiler
	find ./contracts/google -type f | grep -v '\.proto' | xargs rm
	find ./contracts/google -type f -name '*test*.proto' | xargs rm
	rm -rf ./contracts/google/protobuf/util 
	find ./contracts/google -type d -empty -delete
	rm -rf ./tmp
.PHONY: get_known_types

generate_grpc_flights_lib:
	cd contracts && \
	protoc \
		--go_out=../genlibs \
		--go_opt=paths=source_relative \
		--go-grpc_out=../genlibs \
		--go-grpc_opt=paths=source_relative \
		lib/*.proto \
		flights/*.proto
.PHONY: generate_grpc_flights_lib

generate_grpc_trains_lib:
	cd contracts && \
	protoc \
		--go_out=../genlibs \
		--go_opt=paths=source_relative \
		--go-grpc_out=../genlibs \
		--go-grpc_opt=paths=source_relative \
		lib/*.proto \
		trains/*.proto
.PHONY: generate_grpc_trains_lib

flights_server:
	cd flights; go build -o ../flights_server ../flights/cmd/server
.PHONY: flights_server

trains_server:
	cd trains; go build -o ../trains_server ../trains/cmd/server
.PHONY: trains_server

generate_trains_cert:
	# create CA for signing client:
	openssl req -newkey rsa:2048 \
		-new -nodes -x509 \
		-days 3650 \
		-out ./trains/certs/ca.crt \
		-keyout ./trains/certs/ca.key \
		-config ./trains/certs/openssl.cnf 
	# create server private key
	openssl genrsa -out ./trains/certs/server.key 2048
	# create a signing request
	openssl req -new -key \
		./trains/certs/server.key \
		-out ./trains/certs/server.csr \
		-config ./trains/certs/openssl.cnf
	# create SSL certificate:
	openssl x509 -req -in ./trains/certs/server.csr \
		-extfile ./trains/certs/cert.cnf \
		-CA ./trains/certs/ca.crt \
	    -CAkey ./trains/certs/ca.key \
		-out ./trains/certs/server.cert \
		-days 3650 -sha256 -CAcreateserial
.PHONY: generate_trains_cert


flights_docker_image:
	cd flights && docker build -t krakend/grpc_flights_fake:latest .
.PHONY: flights_docker_image

trains_docker_image:
	cd trains && docker build -t krakend/grpc_trains_fake:latest .
.PHONY: trains_docker_image

bin_proto_multiple_files:
	cd contracts && bash ./compile.sh
.PHONY: defs

bin_proto_single_file:
	mkdir -p ./defs
	cd contracts && \
		protoc \
		--descriptor_set_out=../fullcatalog.pb \
		$$(find . -name '*.proto')
.PHONY: fullcatalog

all_images: generate_grpc_flights_lib generate_grpc_trains_lib flights_docker_image trains_docker_image

.PHONY: all_images
