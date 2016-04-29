package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/kelseyhightower/ticker/quote"
	pb "github.com/kelseyhightower/ticker/quotepb"
)

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.FormValue("symbol")
	stockQuote := quote.Quote(symbol)

	data, err := json.MarshalIndent(&stockQuote, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func QuoteProcessorHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var sq pb.Quotes
	err = proto.Unmarshal(data, &sq)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
