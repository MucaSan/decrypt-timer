package service

import (
	"context"
	grpc "decrypter-timer/adapter/grpc/pb"
)

func NewEncryptService() *EncryptService {
	return &EncryptService{}
}

func (ds *EncryptService) EncryptData(ctx context.Context, in *grpc.EncryptRequest) (*grpc.EncryptResponse, error) {
	return nil, nil
}

type EncryptService struct {
	grpc.UnimplementedEncryptServiceServer
}
