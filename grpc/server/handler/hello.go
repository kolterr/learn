package handler

import (
	"context"
	"github.com/kolterr/learn/grpc/server/pb"
)

type Hello struct {
}

func NewHelloHandler() *Hello {
	return &Hello{}
}

func (h *Hello) Say(ctx context.Context, req *pb.SayRequest) (*pb.SayResponse, error) {
	return &pb.SayResponse{Who: "Kolter"}, nil
}
