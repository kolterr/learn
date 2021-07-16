package main

import (
	"github.com/kolterr/learn/grpc/server/handler"
	"github.com/kolterr/learn/grpc/server/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterHelloServer(srv, handler.NewHelloHandler())
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
