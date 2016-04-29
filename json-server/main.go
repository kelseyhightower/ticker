package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/quote", QuoteProcessorHandler).Methods("POST")

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.ListenAndServe(":10000", r)
}
