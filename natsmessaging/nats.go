package natsmessaging

import (
	stan "github.com/nats-io/stan.go"
	"os"
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

func GetSc() *SC {
	return defaultSc
}
