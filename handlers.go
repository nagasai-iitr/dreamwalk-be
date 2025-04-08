package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type FundRequest struct {
	MinAmount       float64 `json:"minAmount"`
	MaxAmount       float64 `json:"maxAmount"`
	ReceiverAddress string  `json:"receiverAddress"`
}

func CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	var req FundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	contract := Entry{
		ID:              uuid.New().String(),
		MinAmount:       req.MinAmount,
		MaxAmount:       req.MaxAmount,
		ReceiverAddress: req.ReceiverAddress,
		Status:          "INIT",
	}

	if err := InsertEntry(contract); err != nil {
		http.Error(w, "Error inserting contract", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contract)
}

func TriggerFundsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := GetInitEntries()
	if err != nil {
		http.Error(w, "Error fetching entries", http.StatusInternalServerError)
		return
	}

	for _, entry := range entries {
		go func(e Entry) {
			//result := triggerAnchorTransfer(entry.MinAmount, entry.MaxAmount, entry.ReceiverAddress)
			result := simulateSmartContractCall(entry.MinAmount, entry.ReceiverAddress)

			status := "FAILED"
			if result {
				status = "SUCCESS"
			}
			UpdateEntryStatus(e.ID, status)
		}(entry)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Triggered fund transfers for INIT entries"))
}

func simulateSmartContractCall(amount float64, address string) bool {
	time.Sleep(1 * time.Second)
	return rand.Intn(2) == 1 // randomly return true or false
}
