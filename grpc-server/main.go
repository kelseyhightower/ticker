package main

import (
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/kelseyhightower/ticker/quotepb"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Quote(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	q := Quote(request.Symbol)
	return &pb.Response{Quote: &q}, nil
}

func (s *server) QuoteStream(request *pb.Request, stream pb.Quoter_QuoteStreamServer) error {
	for {
		q := Quote(request.Symbol)
		err := stream.Send(&pb.Response{Quote: &q})
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Println("Starting Ticker service...")
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuoterServer(s, &server{})
	log.Println("Ticker service started successfully.")
	go s.Serve(ln)

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	log.Fatal(http.ListenAndServe(":10001", nil))
}
