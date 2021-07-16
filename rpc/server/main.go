package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type S struct {
}

type Reply struct {
	Msg []byte
}

func (s *S) Recv(nul *struct{}, reply *Reply) error {
	*reply = Reply{[]byte("Hello world!!!")}
	return nil
}

type S1 struct {
}


func (s *S1) Recv(nul *struct{}, reply *Reply) error {
	*reply = Reply{[]byte("Hello world123113!!!")}
	return nil
}

func main() {
	if err := rpc.Register(new(S)); err != nil {
		panic(err)
	}
	if err := rpc.Register(new(S1)); err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", "127.0.0.1:2015")
	if err != nil {
		fmt.Println("failed to listen tcp:", err)
	}
	defer listen.Close()
	rpc.Accept(listen)
}
