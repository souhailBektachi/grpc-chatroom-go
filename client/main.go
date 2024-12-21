package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/souhailBektachi/grpcWithGo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	fmt.Print("Enter your username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatStreamClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("failed to call Chat: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("failed to receive a message: %v", err)
			}
			fmt.Printf("\n%s :%s \n", msg.User, msg.Message)

		}
	}()

	for {
		scanner.Scan()

		msg := &pb.ChatMessage{
			Message: scanner.Text(),
		}
		if err := stream.Send(&pb.ChatMessage{
			User:    username,
			Message: msg.Message,
		}); err != nil {
			log.Fatalf("failed to send a message: %v", err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

}
