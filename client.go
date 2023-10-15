package main

import (
	"AdvancedChat/chatserver"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter Server IP:Port ::: ")
	reader := bufio.NewReader(os.Stdin)
	serverID, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Failed to read from console :: %v", err)
	}
	serverID = strings.Trim(serverID, "\r\n")

	log.Println("Connecting : " + serverID)

	conn, err := grpc.Dial(serverID, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Faile to conncet to gRPC server :: %v", err)
	}
	defer conn.Close()

	client := chatserver.NewServicesClient(conn)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Faile to conncet to gRPC server :: %v", err)
	}

	ch := clientHandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	bl := make(chan bool)
	<-bl

}

type clientHandle struct {
	stream     chatserver.Services_ChatServiceClient
	clientName string
}

func (ch *clientHandle) clientConfig() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	ch.clientName = strings.Trim(name, "\r\n")

}

func (ch *clientHandle) sendMessage() {
	for {
		reader := bufio.NewReader(os.Stdin)

		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &chatserver.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to the server :: %v", err)
		}
	}
}

func (ch *clientHandle) receiveMessage() {
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		fmt.Printf("%s : %s \n", mssg.Name, mssg.Body)

	}
}
