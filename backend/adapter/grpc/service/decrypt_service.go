package service

import (
	"context"
	pb "decrypter-timer/adapter/grpc/pb"
)

func NewDecryptService() *DecryptService {
	return &DecryptService{}
}

func (ds *DecryptService) DecryptData(ctx context.Context, in *pb.DecryptRequest) (*pb.DecryptResponse, error) {
	return nil, nil
}

type DecryptService struct {
	pb.UnimplementedDecryptServiceServer
}
