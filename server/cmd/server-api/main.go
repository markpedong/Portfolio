package main

import (
	"log"
	"os"
	"portfolio/server-api/handler"
	"portfolio/server-grpc/pb"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(os.Getenv("GRPC_SVC_ADDR"), opts...)
	if err != nil {
		log.Fatal("failed to connect to grpc server: ", err)
	}
	defer conn.Close()

	client := pb.NewApiServiceClient(conn)
	hdl := handler.NewHandler(client, os.Getenv("SECRET_KEY"))
	handler.CreateRoutes(hdl)
	handler.Start(os.Getenv("PORT"))
}
