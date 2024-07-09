package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "itmrchow/go-project/try/grpc/proto"
)

// set port
var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedPokerServer
}

func (s *server) GetNuts(ctx context.Context, req *pb.GetNutsRequest) (*pb.GetNutsResponse, error) {

	println("server get msg")
	fmt.Printf("Hand:+%+v\n", req.Hand)
	fmt.Printf("River:+%+v\n", req.River)

	return &pb.GetNutsResponse{
		Card: []string{"testtest", "testtestyyyy"},
	}, nil
}

func main() {
	flag.Parse() // 解析命令列參數

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPokerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
