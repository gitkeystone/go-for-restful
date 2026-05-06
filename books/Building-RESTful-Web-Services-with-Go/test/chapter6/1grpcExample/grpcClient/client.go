package main

import (
	"context"
	"log"
	pb "protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewMoneyTransactionClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Make a server request.
	r, err := client.MakeTransaction(context.Background(),
		&pb.TransactionRequest{
			From:   from,
			To:     to,
			Amount: amount,
		})
	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)

}
