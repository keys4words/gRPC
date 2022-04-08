package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/keys4words/gRPC/calculator/calcpb"
	"google.golang.org/grpc"
)

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Calc Unary RPC...")

	req := &calcpb.SumRequest{
		FirstNumber:  40,
		SecondNumber: 44,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Calc: %v", res.SumResult)
}

func doServerStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do ServerStreaming RPC...")

	req := &calcpb.PrimeNumberDecompositionRequest{
		Number: 126544848456,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecompositionRequest RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func main() {
	fmt.Println("Calc client started...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calcpb.NewCalculatorServiceClient(conn)

	// doUnary(c)
	doServerStreaming(c)
}
