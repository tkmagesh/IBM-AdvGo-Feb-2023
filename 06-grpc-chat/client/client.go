package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"grpc-chat/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

var client proto.BroadcastClient

func main() {
	timestamp := time.Now()
	name := flag.String("N", "Anon", "The name of the user")
	flag.Parse()
	id := sha256.Sum256([]byte(timestamp.String() + *name))

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client = proto.NewBroadcastClient(conn)
	user := &proto.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}
	connect(user)

	//get input from stdin and send to the server
	done := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msgText := scanner.Text()
			if msgText == "exit" {
				break
			}
			msg := &proto.Message{
				Id:      user.Id,
				Content: msgText,
			}
			if _, err := client.BroadcastMessage(context.Background(), msg); err != nil {
				log.Fatalln(err)
				break
			}
		}
		close(done)
	}()
	<-done

	/* var input string
	fmt.Scanln(&input)
	go func() {
		for i := 0; i < 1000000; i++ {
			msg := &proto.Message{
				Id:      user.Id,
				Content: fmt.Sprintf("Hi [from %q]", *name),
			}
			if _, err := client.BroadcastMessage(context.Background(), msg); err != nil {
				log.Fatalln(err)
				break
			}
		}
		close(done)
	}()
	<-done */
}

func connect(user *proto.User) {
	stream, err := client.SignIn(context.Background(), &proto.Connect{User: user, Active: true})
	if err != nil {
		log.Fatalln(err)
	}
	go func(strm proto.Broadcast_SignInClient) {
		for {
			msg, err := strm.Recv()
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(msg.GetContent())
		}
	}(stream)
}
