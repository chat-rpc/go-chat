// go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"

	"google.golang.org/grpc"
	pb "github.com/chat-rpc/go-chat/demo/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port	= ":50051"
	SVR_NAME= "NBS-CHAT-SERVER"
)

//Server type
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply,error) {
	if in != nil {
		log.Printf("Client :%s visit.",in.Name)
	}
	return &pb.HelloReply{Message:"Hello " + in.Name} , nil
}

func main() {
	//优雅退出监听
	fmt.Println(SVR_NAME," starting...")
	lis, err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("failed to listen: %v" ,err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to save: %v", err)
	}

}

//退出函数
func ExitFunc() {
	fmt.Println("正在退出...")
	os.Exit(0)
}
