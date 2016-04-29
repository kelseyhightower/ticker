package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/kelseyhightower/ticker/quote"
	pb "github.com/kelseyhightower/ticker/quotepb"
)

func protoBody() ([]byte, error) {
	var sq pb.Quotes

	for i := 0; i <= 1000; i++ {
		sq.StockQuotes = append(sq.StockQuotes, &pb.StockQuote{
			Name:              "Company X",
			Symbol:            "x",
			LastPrice:         524.49,
			Change:            15.6,
			ChangePercent:     3.06549549018453,
			Timestamp:         time.Now().Format(time.RFC3339),
			MarketCap:         476497591530,
			Volume:            397562,
			Change_YTD:        532.1729,
			ChangePercent_YTD: -1.44368493773359,
			High:              52499,
			Low:               519.175,
			Open:              519.175,
		})
	}

	return proto.Marshal(&sq)
}

func jsonBody() ([]byte, error) {
	var sq quote.Quotes

	for i := 0; i <= 1000; i++ {
		sq.StockQuotes = append(sq.StockQuotes, &quote.StockQuote{
			Name:             "Company X",
			Symbol:           "x",
			LastPrice:        524.49,
			Change:           15.6,
			ChangePercent:    3.06549549018453,
			Timestamp:        time.Now().Format(time.RFC3339),
			MarketCap:        476497591530,
			Volume:           397562,
			ChangeYTD:        532.1729,
			ChangePercentYTD: -1.44368493773359,
			High:             52499,
			Low:              519.175,
			Open:             519.175,
		})
	}

	return json.Marshal(&sq)
}

func main() {
	log.SetFlags(0)

	format := flag.String("format", "json", "The encoding to use. json or proto")
	duration := flag.Int("duration", 60, "Duration to run in seconds.")
	flag.Parse()

	var data []byte
	var err error

	switch *format {
	case "json":
		data, err = jsonBody()
	case "proto":
		data, err = protoBody()
	}

	if err != nil {
		log.Fatalf("Failed to create request body: %v", err)
	}

	log.Printf("Sending requests for %d seconds using %s encoding. Request size: %d\n", *duration, *format, len(data))

	requestTotal := 0
	go func() {
		for {
			resp, err := http.Post("http://127.0.0.1:10000/quote", "", bytes.NewReader(data))
			if err != nil {
				log.Fatalf("Failed http request: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Error response code: %v", resp.StatusCode)
			}
			requestTotal++
		}
	}()

	time.Sleep(time.Duration(*duration) * time.Second)
	log.Printf("Request Total: %d\n", requestTotal)
}
