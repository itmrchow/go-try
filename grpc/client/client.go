package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "itmrchow/go-project/try/grpc/proto"
)

var (
	// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
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

	printNuts(client, &pb.GetNutsRequest{})

}

func printNuts(client pb.PokerClient, getNutsRequest *pb.GetNutsRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := client.GetNuts(ctx, getNutsRequest)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}

	log.Println(feature)
}
