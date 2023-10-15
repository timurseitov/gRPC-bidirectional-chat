package chatserver

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type messageUnit struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageHandler struct {
	Mque []messageUnit
	mu   sync.Mutex
}

var messageHandleObject = messageHandler{}

type ChatServer struct {
}

func (is *ChatServer) ChatService(csi Services_ChatServiceServer) error {

	clientUniqueCode := rand.Int()
	errch := make(chan error)

	//receive message
	go receiveFromStream(csi, clientUniqueCode, errch)

	//send message
	go sendToStream(csi, clientUniqueCode, errch)

	return <-errch
}

func receiveFromStream(csi_ Services_ChatServiceServer, clienUniqueCode_ int, errch_ chan error) {

	for {
		mssg, err := csi_.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client", err)
			errch_ <- err
		} else {
			messageHandleObject.mu.Lock()

			messageHandleObject.Mque = append(messageHandleObject.Mque, messageUnit{
				ClientName:        mssg.Name,
				MessageBody:       mssg.Body,
				MessageUniqueCode: rand.Int(),
				ClientUniqueCode:  clienUniqueCode_,
			})

			messageHandleObject.mu.Unlock()

			log.Printf("%v", messageHandleObject.Mque[len(messageHandleObject.Mque)-1])
		}
	}
}

func sendToStream(csi_ Services_ChatServiceServer, clientUniqueCode_ int, errch_ chan error) {
	for {
		for {
			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.Mque) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			senderUniqueCode := messageHandleObject.Mque[0].ClientUniqueCode
			senderNameForClient := messageHandleObject.Mque[0].ClientName
			messageForClient := messageHandleObject.Mque[0].MessageBody

			messageHandleObject.mu.Unlock()

			if senderUniqueCode != clientUniqueCode_ {

				err := csi_.Send(&FromServer{Name: senderNameForClient, Body: messageForClient})

				if err != nil {
					errch_ <- err
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.Mque) > 1 {
					messageHandleObject.Mque = messageHandleObject.Mque[1:]
				} else {
					messageHandleObject.Mque = []messageUnit{}
				}

				messageHandleObject.mu.Unlock()
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
