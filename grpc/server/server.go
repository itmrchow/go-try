package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func (s *server) LotsOfReplies(req *pb.HelloRequest, stream pb.Poker_LotsOfRepliesServer) error {
	greetings := []string{"你好", "Hello", "Bonjour", "Hola"}

	// get metadata
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("md not ok")
	}
	metadataMsg := md.Get("key1")[0]
	println("metadataMsg:" + metadataMsg)

	// stream return error sample
	if req.Name == "Jojo" {
		stream.Send(&pb.HelloResponse{Message: req.Name + ":" + "ohla!! ohla!!"})

		return errors.New("new error")
	}

	// stream return success message 
	for _, gStr := range greetings {
		if err := stream.Send(&pb.HelloResponse{Message: req.Name + ":" + gStr + "!"}); err != nil {
			return err
		}
	}
	return nil
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
