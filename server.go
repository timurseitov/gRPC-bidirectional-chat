package main

import (
	"AdvancedChat/chatserver"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}

	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatal("Could not listen ", err)
		return
	}
	log.Println("Listening on :" + Port)

	grpcserver := grpc.NewServer()

	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcserver, &cs)

	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatal("Failed to start grpc server ", err)
	}

}
