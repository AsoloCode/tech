package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"tech/app/repository"
	resp "tech/app/response"
	"tech/cache"
)

func GetOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := chi.URLParam(r, "id")
		fmt.Println(orderId)
		isValidate := IsValidUUID(orderId)
		fmt.Println(isValidate)
		if !isValidate {
			render.JSON(w, r, resp.Error("invalid id"))
		}
		c := cache.GetCache()
		dataFromCache, err := c.GetOrderFromCache(orderId)
		if err == nil {
			render.JSON(w, r, dataFromCache)
			return
		}
		repo := repository.NewOrderRepo()
		data, err := repo.GetById(orderId)
		if err != nil {
			render.JSON(w, r, err.Error())
			return
		}
		c.AddOrderToCache(orderId, data)
		render.JSON(w, r, data)
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
