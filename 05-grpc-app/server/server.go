package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	fmt.Println("Wait for 8 seconds")
	time.Sleep(8 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("timeout occured")
		return nil, errors.New("timeout occured")
	default:
		x := req.GetX()
		y := req.GetY()
		fmt.Printf("Add request received for x = %d and y = %d\n", x, y)
		result := x + y
		fmt.Printf("Sending Add response with result = %d\n", result)
		res := &proto.AddResponse{
			Result: result,
		}
		return res, nil
	}

}

func main() {
	//hosting the service
	asi := &AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
