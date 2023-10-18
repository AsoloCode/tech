package database

import "fmt"

func MockInsert() error {
	const op = "database.MockInsert"
	query := `
	INSERT INTO orders (track_number, entry, locale, customer_id, delivery_service, shardkey, sm_id, oof_shard, date_created)
VALUES ('WBILMTESTTRACK', 'WBIL', 'cen', 'test', 'meest', '9', 99, '1', '2021-11-26');

INSERT INTO item ("chrt_id", track_number, price, rid, "name", sale, "size", total_price, nm_id, brand, status)
VALUES (9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, "0", 317, 2389212, 'Vivienne Sabo', 202);

INSERT INTO item ("chrt_id", track_number, price, rid, "name", sale, "size", total_price, nm_id, brand, status)
VALUES (8823820, 'NUMBERTRACK', 342, 'bc3123457a764ae0btest', 'Kasandras', 20, "0", 318, 1278101, 'Test Adidas', 202);

INSERT INTO delivery ("name", phone, zip, city, address, region, email, order_uid)
VALUES ('Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com', 'b563feb7b2b84b6test');

INSERT INTO payment ("transaction", request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid)
VALUES ('b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0, 'b563feb7b2b84b6test');

`
	_, err := defaultDB.Exec(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
