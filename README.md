# Remote DB Executor (gRPC)

Client (Hub) → gRPC → Server (Spoke) → Fake DB → Response


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

//check
protoc-gen-go --version
protoc-gen-go-grpc --version

mkdir grpc-demo
cd grpc-demo
go mod init grpc-demo

protoc --go_out=. --go-grpc_out=. db.proto

//update dep
go get google.golang.org/grpc && go get google.golang.org/protobuf/reflect/protoreflect && go mod tidy

//terminal 1
//start gRPC server
go run server.go

//terminal 2
go run client.go

---

Or for pg cred
//terminal 1
//start gRPC server
go run server2.go

//terminal 2
go run client.go