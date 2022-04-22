package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/keys4words/gRPC/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog server started!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// listen to ctrl-c to stop server
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until stop signal is received
	<-ch
	fmt.Println("\nTrying to stop server...")
	s.Stop()
	fmt.Println("Server stopped!!!")
	lis.Close()
}
