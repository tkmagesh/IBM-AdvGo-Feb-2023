package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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

	/*
		fmt.Println("Request & Response")
		doRequestResponse(ctx, client)

		fmt.Println("Server Streaming")
		doServerStreaming(ctx, client)
	*/

	fmt.Println("Client Streaming")
	doClientStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
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

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	fmt.Printf("Sending request for generating prime numbers from %d to %d\n", req.Start, req.End)
	clientStream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All prime numbers received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime Number received : %d\n", res.GetPrimeNo())
	}
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	data := []int32{3, 1, 4, 2, 5}
	clientStream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, val := range data {
		fmt.Printf("Sending %d\n", val)
		time.Sleep(500 * time.Millisecond)
		req := &proto.AverageRequest{
			No: val,
		}
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Average : %d\n", res.GetAverage())
}
