package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

/* Request & Response */
func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	fmt.Println("Wait for 3 seconds")
	time.Sleep(3 * time.Second)
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

/* Server Streaming */
func (asi *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Request received for generating prime number from start = %d to end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Sending Prime Number : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			err := serverStream.Send(res)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

/* client streaming */
func (asi *AppServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			fmt.Println("All the requests are received")
			avg := sum / count
			res := &proto.AverageResponse{
				Average: avg,
			}
			serverStream.SendAndClose(res)
			return nil
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Received No : %d\n", req.GetNo())
		sum += req.GetNo()
		count++
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
