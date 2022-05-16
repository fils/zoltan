package main

import (
	"log"
	"net/http"

	"github.com/fils/zoltan/internal/bothook"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/query", bothook.PostCall).Methods("POST")

	log.Println("Starting ON 6789")
	log.Fatal(http.ListenAndServe(":6789", router))
}
