package main

import (
	pb "decrypter-timer/adapter/grpc/pb"
	"decrypter-timer/adapter/grpc/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("\nCouldn't create listener for localhost port :50051. \nError found: %s", err.Error())
	}
	grpcServer := grpc.NewServer()
	decryptService := service.NewDecryptService()
	pb.RegisterDecryptServiceServer(grpcServer, decryptService)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("\nCouldn't serve with provided listener.\nError found: %s", err.Error())
	}
}
