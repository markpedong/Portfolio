package main

import (
	"log"
	"net"
	"os"
	"portfolio/db"
	"portfolio/server-grpc/pb"
	"portfolio/server-grpc/server"
	"portfolio/server-grpc/storer"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error creating database: ", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")

	st := storer.NewPSQLStorer(db.GetDB())
	srv := server.NewServer(st)

	grpcSrv := grpc.NewServer()
	pb.RegisterApiServiceServer(grpcSrv, srv)

	listener, err := net.Listen("tcp", os.Getenv("SVC_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", os.Getenv("SVC_ADDR"))
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
