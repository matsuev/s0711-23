package main

import (
	"log"
	"net"
	"s0711-23/internal/proxyproto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	// Create a new gRPC server
	srv := grpc.NewServer()

	// Register the service with the gRPC server
	proxyproto.RegisterCentrifugoProxyServer(srv, &Service{})

	log.Println("Starting gRPC server on :10000")
	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
