package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/YogeshARIVU/go-transaction-engine/internal/model"
	"github.com/YogeshARIVU/go-transaction-engine/internal/service"
)

type Handler struct {
	service *service.TransactionService
}

func NewHandler(s *service.TransactionService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Transaction(w http.ResponseWriter, r *http.Request) {
	var req model.TransactionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Invalid request"})
		return
	}

	resp := h.service.Process(req)
	writeJSON(w, resp)
}

func (h *Handler) Balance(w http.ResponseWriter, r *http.Request) {
	card := strings.TrimPrefix(r.URL.Path, "/api/card/balance/")

	bal, err := h.service.GetBalance(card)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]float64{"balance": bal})
}

func (h *Handler) Transactions(w http.ResponseWriter, r *http.Request) {
	card := strings.TrimPrefix(r.URL.Path, "/api/card/transactions/")
	writeJSON(w, h.service.GetTransactions(card))
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
