package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	sumpb2 "server/sumpb"

	"google.golang.org/grpc"
)

type server struct{}

func httpServer() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("http listen on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
		return
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called httpHandler...")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to liesten: %v", err)
	}

	go httpServer()

	s := grpc.NewServer()
	sumpb2.RegisterSumServer(s, &server{})

	fmt.Println("grpc listen on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}

// Add returns sum of two integers
func (*server) Add(ctx context.Context, req *sumpb2.SumRequest) (*sumpb2.SumResponse, error) {
	a, b := req.GetNumbers().GetA(), req.GetNumbers().GetB()
	sum := a + b
	return &sumpb2.SumResponse{Result: sum}, nil
}
