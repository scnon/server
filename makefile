
server:
	go build -o ./bin/server main.go

clean:
	rm -rf ./bin

install:
	go build -o ${GOPATH}/bin/server main.go

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protos:
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative proto/*.proto