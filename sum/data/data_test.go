package data

import (
	"testing"
)

func Test_aa(t *testing.T) {
	t.Fatal("haha")
}

func Test_Grpc_Etcd(t *testing.T) {
	/*
		r := discov.NewResolver("localhost:2379")
		resolver.Register(r)

		conn, err := grpc.Dial(r.Scheme()+"://data_goserv",
			grpc.WithBalancerName("round_robin"),
			grpc.WithInsecure())
		if err != nil {
			t.Fatalf(err.Error())
			return
		}

		c := sumpb.NewSumClient(conn)

		// numbers to add
		num := sumpb.Numbers{
			A: 10,
			B: 5,
		}

		// call Add service
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		res, err := c.Add(ctx, &sumpb.SumRequest{Numbers: &num})
		if err != nil {
			log.Fatalf("failed to call Add: %v", err)
		}
		fmt.Println(res.Result)
	*/
}
