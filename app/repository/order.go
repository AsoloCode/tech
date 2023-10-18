package repository

import (
	"fmt"
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

	res, err := repo.db.Query(`
		SELECT * FROM orders
		LEFT JOIN delivery ON orders.order_uid = delivery.order_uid
		LEFT JOIN payment ON orders.order_uid = payment.order_uid
     	LEFT JOIN order_items ON orders.order_uid = order_items.order_id
 		LEFT JOIN item ON order_items.item_id = item.id
	    WHERE orders.order_uid = $1
`, id)
	if err != nil {
		log.Fatalf("qwer: %v", err)
	}
	var o model.Order
	res.Next()
	res.Scan(&o.OrderUID, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSig, &o.CustomerID, &o.DeliveryService, &o.ShardKey, &o.SMID, &o.OofShard, &o.DateCreated)
	defer res.Close()
	//var orders []model.Order
	//for res.Next() {
	//	err := res.Scan(&o.OrderUID, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSig, &o.CustomerID, &o.DeliveryService, &o.ShardKey, &o.SMID, &o.OofShard, &o.DateCreated)
	//	if err != nil {
	//		log.Fatal("chlen adolfa")
	//	}
	//	orders = append(orders, o)
	//}
	//if err := res.Err(); err != nil {
	//	return fmt.Errorf("%s: %w", op, err)
	//}
	//fmt.Println("123", orders)
	fmt.Println(o)
	return o, nil
}
