package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/YogeshARIVU/go-transaction-engine/internal/model"
	"github.com/YogeshARIVU/go-transaction-engine/internal/repository"
	"github.com/YogeshARIVU/go-transaction-engine/internal/util"
)

type TransactionService struct {
	repo *repository.InMemoryRepo
}

func NewTransactionService(r *repository.InMemoryRepo) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) Process(req model.TransactionRequest) model.TransactionResponse {

	card, exists := s.repo.Get(req.CardNumber)
	if !exists {
		return fail("05", "Invalid card")
	}

	if card.Status != model.Active {
		return fail("05", "Card blocked")
	}

	if card.PINHash != util.HashPIN(req.Pin) {
		return fail("06", "Invalid PIN")
	}

	if req.Amount <= 0 {
		return fail("07", "Invalid amount")
	}

	switch model.TransactionType(req.Type) {
	case model.Withdraw:
		if card.Balance < req.Amount {
			s.log(req, "FAILED")
			return fail("99", "Insufficient balance")
		}
		card.Balance -= req.Amount

	case model.Topup:
		card.Balance += req.Amount

	default:
		return fail("08", "Invalid transaction type")
	}

	s.log(req, "SUCCESS")

	return model.TransactionResponse{
		Status:   "SUCCESS",
		RespCode: "00",
		Balance:  card.Balance,
	}
}

func (s *TransactionService) GetBalance(cardNumber string) (float64, error) {
	card, exists := s.repo.Get(cardNumber)
	if !exists {
		return 0, errors.New("card not found")
	}
	return card.Balance, nil
}

func (s *TransactionService) GetTransactions(cardNumber string) []model.Transaction {
	return s.repo.FindByCard(cardNumber)
}

func (s *TransactionService) log(req model.TransactionRequest, status string) {
	tx := model.Transaction{
		ID:         fmt.Sprintf("%d", time.Now().UnixNano()),
		CardNumber: req.CardNumber,
		Type:       model.TransactionType(req.Type),
		Amount:     req.Amount,
		Status:     status,
		Timestamp:  time.Now(),
	}
	s.repo.SaveTransaction(tx)
}

func fail(code, msg string) model.TransactionResponse {
	return model.TransactionResponse{
		Status:   "FAILED",
		RespCode: code,
		Message:  msg,
	}
}
