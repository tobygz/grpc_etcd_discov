package data

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	discov "server/common/discov/etcd"
	sumpb2 "server/sumpb"
)

func Serv() {
	defaultPort := 5000
	var lis net.Listener
	var err error
	for {
		lis, err = net.Listen("tcp", fmt.Sprintf(":%d", defaultPort))
		if err != nil {
			//log.Println("failed to liesten: %v", err)
			defaultPort++
		} else {
			break
		}
	}

	s := grpc.NewServer()
	serv := &server{}
	serv.Init()
	sumpb2.RegisterSumServer(s, serv)

	err = discov.Register(":2379", "dataserv", fmt.Sprintf("127.0.0.1:%d", defaultPort), 5)
	if err != nil {
		return
	}

	fmt.Println("grpc listen on ", defaultPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}
