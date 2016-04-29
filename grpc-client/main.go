package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	pb "github.com/kelseyhightower/ticker/quotepb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var (
		server = flag.String("server", "127.0.0.1:10000", "Server address.")
		symbol = flag.String("symbol", "", "The stock symbol to use.")
	)
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewQuoterClient(conn)

	stream, err := c.QuoteStream(context.Background(), &pb.Request{Symbol: *symbol})
	if err != nil {
		log.Fatal(err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.Quote)
	}
}
