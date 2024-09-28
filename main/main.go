package main

import (
	"database/sql"
	"et_test/proto/et_test"
	store2 "et_test/store"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":7002"
)

func main() {
	db, err := sql.Open("sqlite3", "./thumbnails.db")
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
	}
	defer db.Close()
	store := *store2.NewStore(db)
	store.CreateTable()

	grpcServer := grpc.NewServer()
	et_test.RegisterThumbnailServiceServer(grpcServer, NewServer(store))

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
