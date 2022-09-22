package discov

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	sumpb2 "server/sumpb"
	"testing"
)

func Test_def(t *testing.T) {
	t.Fatal("haha")
}

func Test_register(t *testing.T) {
	//Register("127.0.0.1:2379", "testserver_name", "127.0.0.1:50050", 5)
	ch := make(chan int, 1)
	<-ch
}

func Test_grpccall(t *testing.T) {
	r := NewResolver("localhost:2379")
	resolver.Register(r)

	conn, err := grpc.Dial(r.Scheme()+"://author/testserver_name",
		//grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure())
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer conn.Close()
	c := sumpb2.NewSumClient(conn)

	// numbers to add
	num := sumpb2.Numbers{
		A: 10,
		B: 5,
	}

	// call Add service
	res, err := c.Add(context.Background(), &sumpb2.SumRequest{Numbers: &num})
	if err != nil {
		t.Fatalf("failed to call Add: %v", err)
	}
	fmt.Println("result:", res.Result)
}
