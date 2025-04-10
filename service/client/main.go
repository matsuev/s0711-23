package main

import (
	"context"
	"fmt"
	"log"
	"s0711-23/internal/demoapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("127.0.0.1:10000", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := demoapi.NewDemoServiceClient(conn)

	response, err := client.GetUserByID(context.Background(), &demoapi.GetUserRequest{
		Id: 100500,
	})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(response)
	}
}
