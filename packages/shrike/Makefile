startdb:
	./scripts/start-db.sh
init:
	./scripts/initalize-db.sh
destroy:
	./scripts/purge-db.sh
generate:
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