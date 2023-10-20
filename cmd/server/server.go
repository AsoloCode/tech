package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"strconv"
	"tech/app/controller"
	"tech/cache"
	"tech/natsmessaging"
	"tech/pkg/config"
	"tech/platform/database"
)

// Serve ..
func Serve() {
	// Инициализация конфигурации приложения
	appCfg := config.AppCfg()

	// Подключение к базе данных
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("failed database setup. error: %v", err)
	}

	// Создание таблиц в базе данных (если не существуют)
	err := database.CreateTables()
	if err != nil {
		log.Fatalf("failed to crate database tables. error: %v", err)
	}

	// Получение данных из базы данных и обновление кэша
	orderID := "118127bf-3210-477e-adc0-f3427efea459" // Замените на реальный ID заказа
	cache.InitOrderCache()
	orderData, _ := database.GetOrderDataFromDB(orderID)
	if err != nil {
		log.Fatalf("Failed to retrieve order data from the database: %v", err)
	}
	cache.AddOrderToCache(orderID, orderData)

	//err = database.MockInsert()
	//if err != nil {
	//	log.Fatalf("insert mock error. error: %v", err)
	//}

	// Создание маршрутизатора Chi и применение посредников
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// Установка обработчика для маршрута
	router.Get("/{id}", controller.GetOrderById())

	// Запуск HTTP-сервера
	url := appCfg.Host + ":" + strconv.Itoa(appCfg.Port)
	if err := http.ListenAndServe(url, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	fmt.Printf("server started: %s", url)

	// Параметры подключения к NATS Streaming
	clusterID := "NBWDCKVLB4B4R77UEDO2IE7HTNVCXDBPOS32V2AU6IUQ6F3CFF2NVGEW" // Ваш идентификатор кластера
	clientID := "NBWDCKVLB4B4R77UEDO2IE7HTNVCXDBPOS32V2AU6IUQ6F3CFF2NVGEW"  // Ваш идентификатор клиента
	natsURL := "nats://localhost:4222"

	// Подключение к NATS Streaming
	sc, err := natsmessaging.ConnectNatsStreaming(clusterID, clientID, natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		return
	}
	defer sc.Close()

	// Отправка сообщения
	subject := "orders"
	message := []byte("Hello, NATS Streaming!")

	err = natsmessaging.SendMessage(sc, subject, message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
		return
	}

	log.Printf("Message sent successfully")

	// Подписка на сообщения (если нужно)
	subscription, err := natsmessaging.SubscribeToMessages(sc, subject, func(msg *stan.Msg) {
		log.Printf("Received a message: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
		return
	}
	defer subscription.Unsubscribe()

	// Ожидание сообщений (не завершайте программу сразу)
	select {}
}
