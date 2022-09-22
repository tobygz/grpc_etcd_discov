package http_gate

import (
	"context"
	"fmt"
	"log"
	"net/http"
	sumpb2 "server/sumpb"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called httpHandler...")

	c := sumpb2.NewSumClient(connDataServ)

	// numbers to add
	num := sumpb2.Numbers{
		A: 10,
		B: 5,
	}

	// call Add service
	res, err := c.Add(context.Background(), &sumpb2.SumRequest{Numbers: &num})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}
	fmt.Println(res.Result)
	fmt.Fprintf(w, "sum:%d", res.Result)
}
