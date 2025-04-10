package main

import (
	"log"
	"net"
	"s0711-23/internal/demoapi"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	srv := grpc.NewServer()

	svc := NewMySuperServer()

	demoapi.RegisterDemoServiceServer(srv, svc)

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
