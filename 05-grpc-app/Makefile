generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto

client:
	go run client/client.go

server:
	go run server/server.go

.PHONY: generate client server