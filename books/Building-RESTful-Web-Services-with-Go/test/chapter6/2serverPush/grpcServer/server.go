package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "protofiles"
	"time"
)

const (
	port      = ":50051"
	noOfSteps = 3
)

// server is used to create MoneyTransactionServer
type server struct {
	pb.UnimplementedMoneyTransactionServer
}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer....")
	log.Printf("Amount: $%f, From A/c: %s, To A/c: %s", in.Amount, in.From, in.To)

	// Send streams here
	for i := 0; i < noOfSteps; i++ {
		time.Sleep(2e9)

		// Once task is done, send the successful message
		// back to the client
		err := stream.Send(&pb.TransactionResponse{
			Status:      "good",
			Step:        int32(i),
			Description: fmt.Sprintf("Performing step %d", int32(i)),
		})
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}

	log.Printf("Successfully transferred amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
