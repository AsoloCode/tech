package database

import (
	"fmt"
)

func CreateTables() error {
	const op = "database.CreateTables"

	//	stmt, err := defaultDB.Prepare(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	//	_, err = stmt.Exec()
	//	if err != nil {
	//		return fmt.Errorf("%s", err)
	//}

	query := `
CREATE TABLE IF NOT EXISTS orders (
  order_uid UUID PRIMARY KEY NOT NULL,
  track_number TEXT NOT NULL,
  entry TEXT NOT NULL,
  locale TEXT NOT NULL,
  internal_signature TEXT DEFAULT '' NOT NULL,
  customer_id TEXT NOT NULL,
  delivery_service TEXT NOT NULL,
  shardkey TEXT NOT NULL,
  sm_id INT NOT NULL,
  oof_shard TEXT NOT NULL,
  date_created TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS delivery (
  id SERIAL PRIMARY KEY NOT NULL,
  "name" TEXT NOT NULL,
  phone TEXT NOT NULL,
  zip TEXT NOT NULL,
  city TEXT NOT NULL,
  address TEXT NOT NULL,
  region TEXT NOT NULL,
  email TEXT NOT NULL,
  order_uid UUID NOT NULL,
  CONSTRAINT fk_delivery_order FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);

CREATE TABLE IF NOT EXISTS payment (
  id SERIAL PRIMARY KEY NOT NULL,
  "transaction" TEXT NOT NULL,
  request_id TEXT NOT NULL,
  currency TEXT NOT NULL,
  provider TEXT NOT NULL,
  amount INT NOT NULL,
  payment_dt TIMESTAMP NOT NULL,
  bank TEXT NOT NULL,
  delivery_cost INT NOT NULL,
  goods_total INT NOT NULL,
  custom_fee INT NOT NULL,
  order_uid UUID NOT NULL,
  CONSTRAINT fk_payment_order FOREIGN KEY (order_uid) REFERENCES orders (order_uid)                             
);

CREATE TABLE IF NOT EXISTS item (
  id SERIAL PRIMARY KEY,
  chrt_id INT NOT NULL,
  track_number TEXT NOT NULL,
  price NUMERIC NOT NULL,
  rid TEXT NOT NULL,
  "name" TEXT NOT NULL,
  sale NUMERIC NOT NULL,
  "size" TEXT NOT NULL,
  total_price NUMERIC NOT NULL,
  nm_id INT NOT NULL,
  brand TEXT NOT NULL,
  status INT NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
  item_id INTEGER REFERENCES item (id),
  order_id UUID REFERENCES orders (order_uid),
  CONSTRAINT order_items_pkey PRIMARY KEY (item_id, order_id)
);
`

	_, err := defaultDB.Exec(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
