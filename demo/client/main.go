//client
package main

import (
	"context"
	"log"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/chat-rpc/go-chat/demo/helloworld"
)

const(
	address		= "localhost:50051"
	defaultName	= "worldGo"
)

func main() {
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	fmt.Println("client start...")
	if err != nil {
		log.Fatalf("did not connect: %v",err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	//Contact the server and print out its respone.
	name := defaultName
	
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx,cancel := context.WithTimeout(context.Background(),time.Second)

	defer cancel()
	
	r,err := c.SayHello(ctx,&pb.HelloRequest{Name:name})

	if err != nil {
		log.Fatalf("Coun't greet :%v",err)
	}
	log.Printf("Greeting: %s",r.Message)
	
}
