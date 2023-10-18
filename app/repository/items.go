package repository

import (
	"fmt"
	"tech/platform/database"
)

type ItemRepo struct {
	db *database.DB
}

func NewItemRepo(db *database.DB) (ItemRepository, error) {
	const op = "ItemRepo.NewItemRepo"

	return &ItemRepo{db}, nil
}
func (repo *ItemRepo) GetById(id string) {
	fmt.Println(id)
}
