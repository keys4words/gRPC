package main

import (
	"context"
	"fmt"
	"log"

	"github.com/keys4words/gRPC/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calc client started...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calcpb.NewCalculatorServiceClient(conn)

	doUnary(c)
}

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
