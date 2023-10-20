package natsmessaging

import (
	stan "github.com/nats-io/stan.go"
)

// ConnectNatsStreaming устанавливает соединение с NATS Streaming и возвращает соединение.
func ConnectNatsStreaming(clusterID, clientID, natsURL string) (stan.Conn, error) {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
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
