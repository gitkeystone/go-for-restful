package handlers

import (
	"context"
	pb "encrypt-service/protofiles"
	"encrypt-service/utils"
)

// Encrypter holds the information about methods
type Encrypter struct {
	pb.UnimplementedEncrypterServer
}

// Encrypt converts a message into cipher and returns response
func (g *Encrypter) Encrypt(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	base64EncodedCipher, err := utils.EncryptBytes([]byte(in.Key), []byte(in.Message))
	return &pb.Response{Result: base64EncodedCipher}, err
}

// Decrypt converts a cipher into message and returns response
func (g *Encrypter) Decrypt(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	plaintext, err := utils.DecryptBytes([]byte(in.Key), in.Message)
	return &pb.Response{Result: plaintext}, err
}
