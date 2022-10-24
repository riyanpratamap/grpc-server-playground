package main

import (
	"log"
	"net"

	"github.com/riyanpratamap/grpc-server-playground/chat"
	"google.golang.org/grpc"
)

func main() {
	// Create minimum version of grpc service
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Error creating net tcp on 9000: ", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Erro creating grpc server on 9000: ", err)
	}
}
