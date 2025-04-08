package main

type Entry struct {
	ID              string  `json:"id"`
	MinAmount       float64 `json:"minAmount"`
	MaxAmount       float64 `json:"maxAmount"`
	ReceiverAddress string  `json:"receiverAddress"`
	Status          string  `json:"status"`
}
