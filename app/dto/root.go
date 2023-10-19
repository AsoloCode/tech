package dto

import (
	"tech/app/model"
	"time"
)

// Order представляет собой DTO (Data Transfer Object) для информации о заказе.
type Order struct {
	OrderUID        string    `json:"order_uid"`
	TrackNumber     string    `json:"track_number"`
	CustomerID      string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	ShardKey        string    `json:"shardkey"`
	SMID            int       `json:"sm_id"`
	OofShard        string    `json:"oof_shard"`
	DateCreated     time.Time `json:"date_created"`
}

// ToOrder преобразует модель данных заказа в DTO.
func ToOrder(o *model.Order) *Order {
	return &Order{
		OrderUID:        o.OrderUID,
		TrackNumber:     o.TrackNumber,
		CustomerID:      o.CustomerID,
		DeliveryService: o.DeliveryService,
		ShardKey:        o.ShardKey,
		SMID:            o.SMID,
		OofShard:        o.OofShard,
		DateCreated:     o.DateCreated,
	}
}

// ToOrders преобразует срез заказов в срез DTO.
func ToOrders(orders []*model.Order) []*Order {
	res := make([]*Order, len(orders))
	for i, order := range orders {
		res[i] = ToOrder(order)
	}
	return res
}
