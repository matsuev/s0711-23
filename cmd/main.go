package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const connStr = "postgres://appuser:apppass@127.0.0.1:5432/demodb?sslmode=disable"

func main() {
	connCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalln(err)
	}

	connCfg.MaxConns = 50

	dbConn, err := pgxpool.NewWithConfig(context.Background(), connCfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbConn.Close()

	for range 100 {
		go func() {
			if err := dbConn.Ping(context.Background()); err != nil {
				log.Println(err)
			}
		}()
	}

	var s string
	fmt.Scanln(&s)
}
