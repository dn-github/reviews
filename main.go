package main

import (
	"log"
	"net"

	"github.com/dn-github/reviews/pb"
	"github.com/dn-github/reviews/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	service := server.NewReviewsImpl()
	pb.RegisterReviewServiceServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
