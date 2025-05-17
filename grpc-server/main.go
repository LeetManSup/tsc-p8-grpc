package main

import (
	"context"
	"log"
	"net"

	pb "tsc-p8-grpc/proto/profilepb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProfileServiceServer
}

func (s *server) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	return &pb.ProfileResponse{
		Id:    req.Id,
		Name:  "Ivan Ivanov",
		Email: "ivan@example.com",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("cannot listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{})
	log.Println("gRPC сервер запущен на :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal("cannot serve:", err)
	}
}
