package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"local/_132_grpc/pizza/pb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPizzaServiceServer
}

// implement the interface PizzaServiceServer:

func (s *server) GetProduct(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{
		ProductId: in.ProductId,
		Name:      "Pizza",
		Price:     7.99,
	}, nil
}

func (s *server) ListProducts(in *pb.ProductListRequest, stream pb.PizzaService_ListProductsServer) error {
	return nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPizzaServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
