package main

import (
	"fmt"
	"net/rpc"
)

type Reply struct {
	B []byte
	A string
}

type Reply1 struct {
	Msg []byte
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:2015")
	if err != nil {
		fmt.Println("failed to dial", err)
	}

	var reply Reply
	err = client.Call("S.Recv", &struct{}{}, &reply)
	if err != nil {
		fmt.Println("failed to call", err)
	}
	fmt.Println("reply:", string(reply.B))
	var reply1 Reply1
	err = client.Call("S.Recv", &struct{}{}, &reply1)
	if err != nil {
		fmt.Println("failed to call", err)
	}
	fmt.Println("reply1:", string(reply1.Msg))
	client.Close()
}
