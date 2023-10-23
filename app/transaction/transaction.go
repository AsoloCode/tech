package transaction

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"tech/app/model"
	"tech/cache"
	"tech/platform/database"
)

func InsertOrderData(order *model.Order) error {
	db := database.GetDB()
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed to start a database transaction: %v", err)
		return err
	}
	order.OrderUID = uuid.New().String()
	if err := insertOrder(tx, order); err != nil {
		rollbackTransaction(tx)
		return err
	}

	if err := insertDelivery(tx, order); err != nil {
		rollbackTransaction(tx)
		return err
	}

	if err := insertPayment(tx, order); err != nil {
		rollbackTransaction(tx)
		return err
	}

	if err := insertItems(tx, order.Items); err != nil {
		rollbackTransaction(tx)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit the transaction: %v", err)
		return err
	}
	c := cache.GetCache()
	c.AddOrderToCache(order.OrderUID, *order)
	return nil
}

func insertOrder(tx *sql.Tx, order *model.Order) error {
	_, err := tx.Exec(`
    INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, 
        delivery_service, shardkey, sm_id, oof_shard, date_created)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSig,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.OofShard, order.DateCreated)
	if err != nil {
		log.Printf("Failed to insert data into 'orders' table: %v", err)
		return err
	}

	return nil
}

func insertDelivery(tx *sql.Tx, order *model.Order) error {
	_, err := tx.Exec(`
        INSERT INTO delivery ("name", phone, zip, city, address, region, email, order_uid)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email, order.OrderUID)
	if err != nil {
		log.Printf("Failed to insert data into 'delivery' table: %v", err)
		return err
	}
	return nil
}

func insertPayment(tx *sql.Tx, order *model.Order) error {
	_, err := tx.Exec(`
        INSERT INTO payment ("transaction", request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount,
		order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.OrderUID)
	if err != nil {
		log.Printf("Failed to insert data into 'payment' table: %v", err)
		return err
	}
	return nil
}

func insertItems(tx *sql.Tx, items []model.Item) error {
	for _, item := range items {
		_, err := tx.Exec(`
            INSERT INTO item (chrt_id, track_number, price, rid, "name", sale, "size", total_price, nm_id, brand, status)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NMID, item.Brand, item.Status)
		if err != nil {
			log.Printf("Failed to insert data into 'item' table: %v", err)
			return err
		}
	}
	return nil
}

func rollbackTransaction(tx *sql.Tx) {
	if err := tx.Rollback(); err != nil {
		log.Printf("Failed to rollback the transaction: %v", err)
	}
}
