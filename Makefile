proto:
	protoc -I=./pkg/auth/pb --go_out=. auth.proto

server:
	go run cmd/main.go