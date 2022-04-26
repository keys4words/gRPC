package main

import (
	"context"
	"fmt"
	"io"
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
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogID})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update blog
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Incognito",
		Title:    "New edition of Title",
		Content:  "Updated content",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v\n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

	// delete blog
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogID,
	})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v\n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteRes)

	//list blog
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
