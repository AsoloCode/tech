package repository

import (
	"fmt"
	"tech/platform/database"
)

type DeliveryRepo struct {
	db *database.DB
}

func NewDeliveryRepo(db *database.DB) DeliveryRepository {
	const op = "DeliveryRepo.NewDeliveryRepo"

	return &DeliveryRepo{database.GetDB()}
}
func (repo *DeliveryRepo) GetByOrderId(id string) {
	fmt.Println(id)
}
