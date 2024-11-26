generate:
	protoc --go_out=./pkg --go-grpc_out=./pkg internal/proto_contract/auth.proto