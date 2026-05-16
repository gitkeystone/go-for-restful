package main

import (
	"context"
	"encoding/hex"
	pb "encrypt-client/protofiles"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	// Set up a connection to the server
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewEncrypterClient(conn)

	// Make a server request
	/*
					  bytes   hex
		   AES-128    16      32
		   AES-192    24      48
		   AES-256    32      64
	*/
	AESKey, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	// Encrypt
	r, err := client.Encrypt(context.Background(),
		&pb.Request{
			Message: "Hello World",
			Key:     string(AESKey),
		})
	if err != nil {
		log.Fatalf("Could not encrpt: %v", err)
	}
	log.Printf("密文: %x", r.Result)

	// Decrypt
	r, err = client.Decrypt(context.Background(),
		&pb.Request{
			Message: r.Result,
			Key:     string(AESKey),
		})
	if err != nil {
		log.Fatalf("Could not encrpt: %v", err)
	}
	log.Printf("明文: %s", r.Result)
}
