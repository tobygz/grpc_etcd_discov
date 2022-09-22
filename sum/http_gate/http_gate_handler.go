package http_gate

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"server/sum/sumpb"
)


func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called httpHandler...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := sumpb.NewSumClient(conn)

	// numbers to add
	num := sumpb.Numbers{
		A: 10,
		B: 5,
	}

	// call Add service
	res, err := c.Add(context.Background(), &sumpb.SumRequest{Numbers: &num})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}
	fmt.Println(res.Result)
	fmt.Fprintf(w,"sum:%d", res.Result)
}