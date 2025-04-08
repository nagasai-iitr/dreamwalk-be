package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	_ = godotenv.Load()
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/sendFundsEntry", CreateEntryHandler).Methods("POST")
	r.HandleFunc("/triggerFunds", TriggerFundsHandler).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
