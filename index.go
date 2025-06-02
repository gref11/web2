package main

import (
	"log"
	"net/http"
	"net/http/cgi"

	"web3/handlers"
)

func main() {
	handler, err := handlers.NewHandler()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
	defer handler.Storage.Close()

	mux := http.NewServeMux()
	// mux.Handle("/", http.FileServer(http.Dir("./static")))
	// mux.HandleFunc("/", staticHandler)
	mux.HandleFunc("/", handler.MainHandler)
	// err = http.ListenAndServe(":8080", mux)
	err = cgi.Serve(mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
