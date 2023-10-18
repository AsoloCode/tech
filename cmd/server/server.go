package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strconv"
	"tech/app/controller"
	"tech/pkg/config"
	"tech/platform/database"
)

// Serve ..
func Serve() {
	appCfg := config.AppCfg()

	// connect to DB
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("failed database setup. error: %v", err)
	}
	err := database.CreateTables()
	if err != nil {
		log.Fatalf("failed to crate database tables. error: %v", err)
	}

	//err = database.MockInsert()
	//if err != nil {
	//	log.Fatalf("insert mock error. error: %v", err)
	//}

	router := chi.NewRouter()

	router.Use(middleware.URLFormat)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Get("/", controller.GetOrderById())

	// start http server

	url := appCfg.Host + ":" + strconv.Itoa(appCfg.Port)

	if err := http.ListenAndServe(url, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	fmt.Printf("server started: %s", url)
}
