package main

import (
	"fmt"
	"net/http"

	"github.com/YogeshARIVU/go-transaction-engine/internal/handler"
	"github.com/YogeshARIVU/go-transaction-engine/internal/model"
	"github.com/YogeshARIVU/go-transaction-engine/internal/repository"
	"github.com/YogeshARIVU/go-transaction-engine/internal/service"
	"github.com/YogeshARIVU/go-transaction-engine/internal/util"
)

func main() {

	repo := repository.NewInMemoryRepo()

	repo.Save(&model.Card{
		CardNumber: "4123456789012345",
		CardHolder: "John Doe",
		PINHash:    util.HashPIN("1234"),
		Balance:    1000,
		Status:     model.Active,
	})

	svc := service.NewTransactionService(repo)
	h := handler.NewHandler(svc)

	http.HandleFunc("/api/transaction", h.Transaction)
	http.HandleFunc("/api/card/balance/", h.Balance)
	http.HandleFunc("/api/card/transactions/", h.Transactions)

	fmt.Println("🚀 Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
