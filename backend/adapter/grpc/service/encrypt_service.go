package service

import (
	"context"
	pb "decrypter-timer/adapter/grpc/pb"
)

func NewEncryptService() *EncryptService {
	return &EncryptService{}
}

func (ds *EncryptService) EncryptData(ctx context.Context, in *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	return nil, nil
}

type EncryptService struct {
	pb.UnimplementedEncryptServiceServer
}
