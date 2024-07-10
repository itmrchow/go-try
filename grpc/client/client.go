package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "itmrchow/go-project/try/grpc/proto"
)

var (
	// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	userName   = flag.String("name", "Jojo", "client send username")
	// serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // 建立一個不安全的憑證 , 使用未加密TCP

	conn, err := grpc.NewClient(*serverAddr, opts...) //
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPokerClient(conn)

	// printNuts(client, &pb.GetNutsRequest{})

	// sayHello(client, &pb.HelloRequest{Name: *userName})

	sayManyHello(client, []string{"Jeff", "Jojo", "John", "Josh"})

}

func printNuts(client pb.PokerClient, getNutsRequest *pb.GetNutsRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := client.GetNuts(ctx, getNutsRequest)
	if err != nil {
		log.Fatalf("client.GetNuts failed: %v", err)
	}

	log.Println(feature)
}

func sayHello(client pb.PokerClient, helloReq *pb.HelloRequest) {

	// set metadata in context
	md := metadata.Pairs(
		"key1", "hey",
		"key1", "val1-2", // "key1" will have map value []string{"val1", "val1-2"}
		"key2", "val2",
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	send, _ := metadata.FromOutgoingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, md))

	// send request to grpc server
	helloStream, err := client.LotsOfReplies(ctx, helloReq)
	if err != nil {
		log.Fatalf("client.LotsOfReplies failed: %v", err)
	}

	for {
		helloResp, err := helloStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("client.LotsOfReplies failed: err: %v", err)
		} else {
			log.Println(helloResp.Message)
		}
	}
}

func sayManyHello(client pb.PokerClient, names []string) {

	// Create stream
	stream, err := client.LotsOfGreetings(context.Background())
	if err != nil {
		log.Fatalf("client.LotsOfGreetings failed: %v", err)
	}

	// Forloop , send msg in stream
	for _, n := range names {
		if err := stream.Send(&pb.HelloRequest{Name: n}); err != nil {
			log.Fatalf("stream.Send failed: %v", err)
		}
	}

	// close stream & get response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream close failed: %v", err)
	}

	log.Println("ResponseMsg:" + resp.Message)

}
