package main

import (
	"context"
	"log"
	"time"

	pb "tsc-p8-grpc/proto/profilepb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("не удалось подключиться:", err)
	}
	defer conn.Close()

	client := pb.NewProfileServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetProfile(ctx, &pb.ProfileRequest{Id: 1})
	if err != nil {
		log.Fatal("ошибка при вызове:", err)
	}

	log.Printf("Получен профиль: ID=%d, Name=%s, Email=%s", resp.Id, resp.Name, resp.Email)
}
