package data

import (
	"context"
	"server/sum/sumpb"
)

type server struct{
	processer *process
}

func (s *server) Init(){
	s.processer = &process{}
	s.processer.Init()
}

// Add returns sum of two integers
func (s *server) Add(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	return s.processer.Add(ctx,req)
}
