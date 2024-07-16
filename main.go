package main

import (
	"log"
	"net"

	"main/api/proto/server"

	"main/api/service"
	"main/database"

	"google.golang.org/grpc"
)

func main() {
	err := database.SetUp()
	if err != nil {
		log.Fatalf("failed to database: %v", err)
		panic(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err.Error())
	}

	s := grpc.NewServer()
	server.RegisterUserServiceServer(s, &service.UserService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err.Error())
	}
}
