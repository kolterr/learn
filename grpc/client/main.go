package main

import (
	"context"
	"fmt"
	"github.com/kolterr/learn/grpc/client/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	reply, err := client.Say(context.Background(), &pb.SayRequest{Name: "WangYang"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Who)
}
