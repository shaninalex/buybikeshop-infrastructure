package main

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewCatalogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ProductList(ctx, &pb.ProductListRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, product := range r.Products {
		fmt.Println(product.Id, product.Title, product.Description, product.ShortDescription, product.CreatedAt.AsTime())
	}
}
