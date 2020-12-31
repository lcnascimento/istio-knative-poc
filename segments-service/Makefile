GOPATH=$(shell go env GOPATH)
PATH=$(GOPATH)/bin:$(shell echo $$PATH)

@:
	@go run cmd/grpcserver/main.go

protoc:
	@echo "compiling ./application/grpc/proto/*.proto into ./application/grpc/proto/*.pb.go ..."
	@protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--experimental_allow_proto3_optional \
	  application/grpc/proto/*.proto
