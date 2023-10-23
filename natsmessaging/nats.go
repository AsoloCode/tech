package natsmessaging

import (
	_ "database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	stan "github.com/nats-io/stan.go"
	"log"
	"os"
	"tech/app/model"
	"tech/app/transaction"
)

type SC struct{ stan.Conn }

// database instance
var defaultSc = &SC{}

// ConnectNatsStreaming устанавливает соединение с NATS Streaming и возвращает соединение.
func ConnectNatsStreaming() (stan.Conn, error) {
	sc, err := stan.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"), stan.NatsURL(os.Getenv("NATS_URL")))
	if err != nil {
		return nil, err
	}
	return sc, nil
}

// SendMessage отправляет сообщение в NATS Streaming.
func SendMessage(sc stan.Conn, subject string, message []byte) error {
	err := sc.Publish(subject, message)
	return err
}

// SubscribeToMessages подписывается на сообщения в NATS Streaming и обрабатывает их.
func SubscribeToMessages(sc stan.Conn, subject string, handler func(msg *stan.Msg)) (stan.Subscription, error) {
	subscription, err := sc.Subscribe(subject, handler)
	return subscription, err
}

// Обработчик сообщений NATS
func HandleNATSMessage(msg *stan.Msg) {
	// Распаковываем сообщение из NATS
	var orderData *model.Order
	if err := json.Unmarshal(msg.Data, &orderData); err != nil {
		log.Printf("Failed to unmarshal NATS message: %v", err)
		return
	}

	// Подключение к базе данных
	db, err := sqlx.Open("your-database-driver", "connection-string")
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return
	}
	defer db.Close()

	// Выполните операции записи данных в базу данных, используя функции из пакета transaction, например:
	if err := transaction.InsertOrderData(&model.Order{}); err != nil {
		log.Printf("Failed to insert data into the database: %v", err)
		return
	}

	// Все успешно обработано
	log.Printf("Data from NATS message inserted into the database")
}

func GetSc() *SC {
	return defaultSc
}
