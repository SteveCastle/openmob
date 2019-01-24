PROTOC_ZIP=protoc-3.3.0-linux-x86_64.zip
startdb:
	./scripts/start-db.sh
init:
	./scripts/initialize-db.sh
destroy:
	./scripts/purge-db.sh
install:
	go get -u github.com/rubenv/sql-migrate
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
generate:
	curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/$(PROTOC_ZIP)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	rm -f $(PROTOC_ZIP)
	./scripts/protogen.sh
build-server:
	cd src/cmd/grpc-server; go build .
build-client:
	cd src/cmd/grpc-client; go build .
build: build-server build-client
start: build
	cd src/cmd/grpc-server;./grpc-server
test-client:
	cd src/cmd/grpc-client;./grpc-client -server=localhost:9090
bootstrap: startdb install init generate build