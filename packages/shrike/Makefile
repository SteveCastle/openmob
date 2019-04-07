OSFLAG 				:=
ifeq ($(OS),Windows_NT)
	OSFLAG += -D WIN32
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		OSFLAG += -D AMD64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		OSFLAG += -D IA32
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		PROTOC_ZIP=protoc-3.3.0-linux-x86_64.zip
		OSFLAG += -D LINUX
	endif
	ifeq ($(UNAME_S),Darwin)
		PROTOC_ZIP=protoc-3.3.0-osx-x86_64.zip
		OSFLAG += -D OSX
	endif
		UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		OSFLAG += -D AMD64
	endif
		ifneq ($(filter %86,$(UNAME_P)),)
	OSFLAG += -D IA32
		endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		OSFLAG += -D ARM
	endif
endif

startdb:
	./scripts/start-db.sh
init:
	./scripts/initialize-db.sh
destroy:
	./scripts/purge-db.sh
install:
	go get -u github.com/rubenv/sql-migrate
	go get -u gnorm.org/gnorm
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	./scripts/install-protoc.sh

codegen:
	gnorm gen
protogen:
	./scripts/protogen.sh
generate: codegen protogen

build-client:
	cd src/cmd/sample-data-client; go build .
build-server:
	cd src/cmd/grpc-server; go build .
build: build-server
start: build-server
	cd src/cmd/grpc-server; ./grpc-server
load-data: build-client
	cd src/cmd/sample-data-client; ./sample-data-client -server=localhost:9090
bootstrap: startdb install init generate build
schema-update: destroy init generate build

swagger:
	./scripts/start-swagger.sh