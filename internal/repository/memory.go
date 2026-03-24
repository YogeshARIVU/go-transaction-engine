package repository

import "github.com/YogeshARIVU/go-transaction-engine/internal/model"

type InMemoryRepo struct {
	cards        map[string]*model.Card
	transactions map[string][]model.Transaction
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		cards:        make(map[string]*model.Card),
		transactions: make(map[string][]model.Transaction),
	}
}

func (r *InMemoryRepo) Get(cardNumber string) (*model.Card, bool) {
	card, ok := r.cards[cardNumber]
	return card, ok
}

func (r *InMemoryRepo) Save(card *model.Card) {
	r.cards[card.CardNumber] = card
}

func (r *InMemoryRepo) SaveTransaction(tx model.Transaction) {
	r.transactions[tx.CardNumber] = append(r.transactions[tx.CardNumber], tx)
}

func (r *InMemoryRepo) FindByCard(cardNumber string) []model.Transaction {
	return r.transactions[cardNumber]
}
