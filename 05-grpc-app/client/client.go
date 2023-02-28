package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	resp, err := client.Add(timeoutCtx, addRequest)
	if err != nil {
		if code := status.Code(err); code == codes.DeadlineExceeded {
			fmt.Println("Timeout occured")
			return
		}
	}
	fmt.Println("Add Response :", resp.GetResult())
}
