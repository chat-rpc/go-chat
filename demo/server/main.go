// go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	pb "github.com/chat-rpc/go-chat/demo/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port	= ":50051"
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
	c := make(chan os.Signal)
	//set signal will exec quit
	signal.Notify(c,syscall.SIGHUP,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT,syscall.SIGUSR1,syscall.SIGUSR2)

	go func() {
		for s := range c {
			switch s {
				case syscall.SIGHUP,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT :
					fmt.Println("退出",s)
					ExitFunc(s)
				case syscall.SIGUSR1,syscall.SIGUSR2:
					ExitFunc(s)
			chatsvr	default:
					ExitFunc(s)

			}
		}
	}

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
func ExitFunc(s int) {
	fmt.Println("quit:",s)
	fmt.Println("正在退出...")
	os.Exit(0)
}
