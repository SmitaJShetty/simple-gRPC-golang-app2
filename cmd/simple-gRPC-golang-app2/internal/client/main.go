package main

import (
	"Vakratunda/simple-gRPC-golang/cmd/rpc"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("app2 dialing in ...")
	addClient()
}

func addClient() {
	conn, connErr := grpc.Dial(":8888", grpc.WithInsecure())
	if connErr != nil {
		log.Fatal("failed to connect on port 8888")
	}
	defer conn.Close()
	if conn == nil {
		log.Fatal("listener was empty")
	}

	c := rpc.NewApp1Client(conn)
	if c == nil {
		log.Fatal("created new api client")
	}

	resp, err := c.SayHi(context.Background(), &rpc.HiRequest{Request: "client hi"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Response)

}
