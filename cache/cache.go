package cache

// Кэширование структуры данных
type OrderCache struct {
	Orders map[string]string // Используйте подходящий тип данных для хранения заказов
}

var orderCache OrderCache // Объявление переменной типа OrderCache

// Инициализация кэша
func InitOrderCache() {
	orderCache.Orders = make(map[string]string)
}

// Добавление данных в кэш
func AddOrderToCache(orderID string, orderData string) {
	orderCache.Orders[orderID] = orderData
}

// Получение данных из кэша
func GetOrderFromCache(orderID string) string {
	return orderCache.Orders[orderID]
}
