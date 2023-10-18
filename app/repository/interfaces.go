package repository

import "tech/app/model"

type OrderRepository interface {
	GetById(id string) (model.Order, error)
}
type PaymentRepository interface {
	GetById(id string)
}
type DeliveryRepository interface {
	GetByOrderId(id string)
}

type ItemRepository interface {
	GetById(id string)
}
