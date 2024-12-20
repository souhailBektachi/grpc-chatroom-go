package main

import (
	"log"
	"net"
	"sync"

	pb "github.com/souhailBektachi/grpcWithGo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type ChatServer struct {
	pb.ChatStreamServer
	clents map[pb.ChatStream_ChatServer]chan *pb.ChatMessage
	mu     sync.Mutex
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := &ChatServer{
		clents: make(map[pb.ChatStream_ChatServer]chan *pb.ChatMessage),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatStreamServer(grpcServer, chatServer)
	log.Printf("Server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *ChatServer) Chat(stream pb.ChatStream_ChatServer) error {
	clientChan := make(chan *pb.ChatMessage)
	s.mu.Lock()
	s.clents[stream] = clientChan
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		delete(s.clents, stream)
		s.mu.Unlock()
		close(clientChan)
		disconectingMsg := &pb.ChatMessage{Message: "Client disconnected"}
		s.mu.Lock()
		for _, ch := range s.clents {
			ch <- disconectingMsg
		}
		s.mu.Unlock()
	}()
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Client disconnected: %v", err)
				s.mu.Lock()
				delete(s.clents, stream)
				s.mu.Unlock()
				return

			}
			log.Printf("Message from client: %s", msg.Message)
			s.mu.Lock()
			for _, ch := range s.clents {
				ch <- msg
			}
			s.mu.Unlock()

		}
	}()

	for msg := range clientChan {
		if err := stream.Send(msg); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
	return nil
}
