serve:
	go run main.go serve

grpc_gen:
	protoc --go_out=./internal/grpc \
		   --go-grpc_out=./internal/grpc \
		   ./internal/grpc/users.proto


