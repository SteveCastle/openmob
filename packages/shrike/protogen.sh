protoc --proto_path=src/api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:src/pkg/api/v1 cause.proto
protoc --proto_path=src/api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:src/pkg/api/v1 cause.proto
protoc --proto_path=src/api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:src/api/swagger/v1 cause.proto