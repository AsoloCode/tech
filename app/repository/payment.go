package repository

import (
	"fmt"
	"tech/platform/database"
)

type PaymentRepo struct {
	db *database.DB
}

func NewPaymentRepo(db *database.DB) (PaymentRepository, error) {
	const op = "PaymentRepo.NewPaymentRepo"

	return &PaymentRepo{db}, nil
}
func (repo *PaymentRepo) GetById(id string) {
	fmt.Println(id)
}
