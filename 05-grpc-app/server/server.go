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
	opCount map[string]int
	proto.UnimplementedAppServiceServer
}

/* Request & Response */
func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	asi.opCount["Add"]++
	fmt.Println("Value from context : ", ctx.Value("key-1"))
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
	asi.opCount["GeneratePrimes"]++
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
	asi.opCount["CalculateAverage"]++
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

/* Bidirectional Streaming */
func (asi *AppServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		asi.opCount["Greet"]++
		personName := *req.GetPerson()
		firstName := personName.GetFirstName()
		lastName := personName.GetLastName()
		fmt.Printf("Received req : %s %s\n", firstName, lastName)
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day", firstName, lastName)
		fmt.Printf("Sending response : %s\n", greetMsg)
		time.Sleep(500 * time.Millisecond)
		res := &proto.GreetResponse{
			Message: greetMsg,
		}
		if err := serverStream.Send(res); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func main() {
	//hosting the service
	asi := &AppServiceImpl{
		opCount: make(map[string]int),
	}
	done := make(chan struct{})
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)

	go func() {
		fmt.Println("Hit ENTER to print stats.... EXIT to shutdown..")
		var input string
		for {
			fmt.Scanln(&input)
			if input == "EXIT" {
				break
			}
			fmt.Println(asi.opCount)
		}
		close(done)
		grpcServer.Stop()
	}()
	go func() {
		grpcServer.Serve(listener)
	}()
	<-done
	fmt.Println("Server shutdown...")
}
