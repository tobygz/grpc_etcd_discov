package data

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/sum/sumpb"
)

func Serv() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to liesten: %v", err)
	}

	s := grpc.NewServer()
	serv := &server{}
	serv.Init()
	sumpb.RegisterSumServer(s, serv)

	//discov.Register(context.Background(), "data_goserv", "127.0.0.1:8080")

	fmt.Println("grpc listen on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}
