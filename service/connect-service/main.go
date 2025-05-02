package main

import (
	"context"
	"log"
	"net"
	"s0711-23/internal/proxyproto"
	"s0711-23/internal/userdb"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

const connString = "postgres://appuser:apppass@127.0.0.1:5432/userdb?sslmode=disable"

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
	}

	store := userdb.New(conn)

	svc := NewService(store)

	// Create a new gRPC server
	srv := grpc.NewServer()

	// Register the service with the gRPC server
	proxyproto.RegisterCentrifugoProxyServer(srv, svc)

	log.Println("Starting gRPC server on :10000")
	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
