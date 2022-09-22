package http_gate

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net/http"
	discov "server/common/discov/etcd"
)

var connDataServ *grpc.ClientConn

func init() {
	r := discov.NewResolver("localhost:2379")
	resolver.Register(r)

	var err error
	connDataServ, err = grpc.Dial(r.Scheme()+"://author/dataserv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func HttpServer() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("http listen on 8081")
	err := discov.Register(":2379", "http_gate", "127.0.0.1:50051", 5)
	if err != nil {
		return
	}
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
