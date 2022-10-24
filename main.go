package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create minimum version of grpc service
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Error creating net tcp on 9000: ", err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Erro creating grpc server on 9000: ", err)
	}
}
