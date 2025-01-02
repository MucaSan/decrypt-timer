package service

import (
	"context"
	grpc "decrypter-timer/adapter/grpc/pb"
)

func NewDecryptService() *DecryptService {
	return &DecryptService{}
}

func (ds *DecryptService) DecryptData(ctx context.Context, in *grpc.DecryptRequest) (*grpc.DecryptResponse, error) {
	return nil, nil
}

type DecryptService struct {
	grpc.UnimplementedDecryptServiceServer
}
