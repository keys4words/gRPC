package main

import (
	"context"
	"fmt"
	"log"

	"github.com/keys4words/gRPC/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client started...")
	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)

	blog := &blogpb.Blog{
		AuthorId: "Maxim",
		Title:    "My First Post",
		Content:  "Content of the first post",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", res)
	blogID := res.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "6266793a51d939d2a24d6e9a"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)
}
