package controller

import (
	"net/http"
	"tech/app/repository"
)

func GetOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := repository.NewOrderRepo()
		_, err := repo.GetById("118127bf-3210-477e-adc0-f3427efea459")
		if err != nil {
			return
		}
	}
}
