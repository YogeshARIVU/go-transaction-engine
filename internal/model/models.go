package model

import "time"

type CardStatus string

const (
	Active  CardStatus = "ACTIVE"
	Blocked CardStatus = "BLOCKED"
)

type Card struct {
	CardNumber string
	CardHolder string
	PINHash    string
	Balance    float64
	Status     CardStatus
}

type TransactionType string

const (
	Withdraw TransactionType = "withdraw"
	Topup    TransactionType = "topup"
)

type Transaction struct {
	ID         string
	CardNumber string
	Type       TransactionType
	Amount     float64
	Status     string
	Timestamp  time.Time
}

type TransactionRequest struct {
	CardNumber string  `json:"cardNumber"`
	Pin        string  `json:"pin"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
}

type TransactionResponse struct {
	Status   string  `json:"status"`
	RespCode string  `json:"respCode"`
	Message  string  `json:"message,omitempty"`
	Balance  float64 `json:"balance,omitempty"`
}
