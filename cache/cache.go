package cache

import (
	"errors"
	"tech/app/model"
)

// Кэширование структуры данных
type OrderCache struct {
	Orders map[string]model.Order
}

var orderCache = &OrderCache{} // Объявление переменной типа OrderCache

// Инициализация кэша
func InitOrderCache() {
	orderCache.Orders = make(map[string]model.Order)
}

// Добавление данных в кэш
func (c *OrderCache) AddOrderToCache(orderID string, orderData model.Order) {
	c.Orders[orderID] = orderData
}

// Получение данных из кэша
func (c *OrderCache) GetOrderFromCache(orderID string) (model.Order, error) {

	if value, ok := c.Orders[orderID]; ok {
		return value, nil
	}
	return model.Order{}, errors.New("Заказ не найден в кеше")
}

func GetCache() *OrderCache {
	return orderCache
}
