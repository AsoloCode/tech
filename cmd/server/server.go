package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"strconv"
	"tech/app/controller"
	"tech/app/model"
	"tech/app/transaction"
	"tech/cache"
	"tech/mypackage/pb"
	"tech/natsmessaging"
	"tech/pkg/config"
	"tech/platform/database"
)

// Serve ..
func Server() {
	// Инициализация конфигурации приложения
	appCfg := config.AppCfg()

	// Подключение к базе данных
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("failed database setup. error: %v", err)
	}

	// Создание таблиц в базе данных (если не существуют)
	err := database.CreateTables()
	if err != nil {
		log.Fatalf("failed to create database tables. error: %v", err)
	}

	cache.InitOrderCache()

	// Создание маршрутизатора Chi и применение посредников
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// Установка обработчика для маршрута
	router.Get("/{id}", controller.GetOrderById())

	// Подключение к NATS Streaming
	sc, err := natsmessaging.ConnectNatsStreaming()
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		return
	}

	subject := "orders"
	order := pb.Order{}
	modelOrder := model.Order{}

	// Подписка на сообщения (если нужно)

	subscription, err := natsmessaging.SubscribeToMessages(sc, subject, func(msg *stan.Msg) {

		if err := proto.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return
		}

		modelOrder.OrderFromProto(&order) // Преобразование protobuf Order в модель Order

		// Сохраняем данные из NATS в базе данных

		err := transaction.InsertOrderData(&modelOrder)
		if err != nil {
			log.Printf("Failed to insert order data into the database: %v", err)
		} else {
			log.Printf("Order data inserted into the database successfully")
		}
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
		return
	}

	defer func() {
		if err := subscription.Unsubscribe(); err != nil {
			log.Printf("Failed to unsubscribe: %v", err)
		}
	}()
	// Запуск HTTP-сервера
	url := appCfg.Host + ":" + strconv.Itoa(appCfg.Port)
	if err := http.ListenAndServe(url, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Printf("Server started: %s", url)

}
