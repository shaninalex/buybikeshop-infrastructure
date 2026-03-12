package main

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The datasource port")
)

type server struct {
	pb.UnimplementedCatalogServer
}

func (c server) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	//TODO implement me
	panic("implement me")
}

func (c server) ProductGet(ctx context.Context, request *pb.ProductGetRequest) (*pb.ProductGetReply, error) {
	//TODO implement me
	panic("implement me")
}

func (c server) ProductVariantList(ctx context.Context, request *pb.ProductVariantListRequest) (*pb.ProductVariantListReply, error) {
	//TODO implement me
	panic("implement me")
}

func (c server) ProductVariantGet(ctx context.Context, request *pb.ProductVariantGetRequest) (*pb.ProductVariantGetReply, error) {
	//TODO implement me
	panic("implement me")
}

func (c server) mustEmbedUnimplementedserver() {
	//TODO implement me
	panic("implement me")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCatalogServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
