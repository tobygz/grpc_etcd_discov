protoc -I. --go_out=plugins=grpc:. sum.proto

#env
#go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
#go install github.com/golang/protobuf/protoc-gen-go