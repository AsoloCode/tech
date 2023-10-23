package repository

import (
	"github.com/google/uuid" // Импорт пакета для работы с UUID
	"log"
	"tech/app/model"
	"tech/platform/database"
)

type OrderRepo struct {
	db *database.DB
}

func NewOrderRepo() OrderRepository {
	const op = "OrderRepo.NewOrderRepo"

	return &OrderRepo{database.GetDB()}
}

func (repo *OrderRepo) GetById(id string) (model.Order, error) {
	const op = "OrderRepo.GetById"

	// Преобразование строки UUID в формат UUID из пакета google/uuid
	orderUID, err := uuid.Parse(id)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}

	res, err := repo.db.Query(`
		SELECT
    o.*,
    d.name, d.phone, d.zip, d.city, d.address, d.region, d.email,
    p.transaction, p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee,
    i.*
		FROM orders AS o
		LEFT JOIN delivery AS d ON o.order_uid = d.order_uid
		LEFT JOIN payment AS p ON o.order_uid = p.order_uid
		LEFT JOIN order_items AS oi ON o.order_uid = oi.order_id
		LEFT JOIN item AS i ON oi.item_id = i.id
		WHERE o.order_uid = $1
	`, orderUID)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	defer res.Close()

	var o model.Order

	for res.Next() {
		var delivery model.Delivery
		var payment model.Payment
		var item model.Item

		if err := res.Scan(
			&o.OrderUID, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSig, &o.CustomerID,
			&o.DeliveryService, &o.ShardKey, &o.SMID, &o.OofShard, &o.DateCreated,
			&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email,
			&payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee,
			&item.NMID, &item.ChrtID, &item.TrackNumber, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NMID, &item.Brand, &item.Status,
		); err != nil {
			log.Fatalf("%s: %v", op, err)
		}

		o.Delivery = delivery
		o.Payment = payment
		o.Items = append(o.Items, item)
	}

	if err := res.Err(); err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	return o, nil
}
