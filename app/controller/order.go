package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"tech/app/repository"
	resp "tech/app/response"
)

func GetOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := chi.URLParam(r, "id")
		isValidate := IsValidUUID(orderId)
		if !isValidate {
			render.JSON(w, r, resp.Error("invalid id"))
		}

		repo := repository.NewOrderRepo()
		data, err := repo.GetById(orderId)
		if err != nil {
			return
		}
		render.JSON(w, r, data)
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
