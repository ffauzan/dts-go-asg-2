package main

import (
	"asg-2/config"
	"asg-2/order"
	"asg-2/storage/sql"
	"asg-2/transport/rest"
	"log"
	"net/http"
)

func main() {
	// Config
	c, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	// Store
	dbUrl := "postgres://" + c.DBUser + ":" + c.DBPassword + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + "?sslmode=disable"

	store, err := sql.New(dbUrl)
	if err != nil {
		log.Fatalf("failed to create store: %v", err)
	}

	// Order
	orderRepo := sql.NewOrderRepo(store)
	orderService := order.NewService(orderRepo)

	// REST
	router := rest.NewRouter(orderService)
	log.Println("Starting server on port:", c.AppPort)
	if err := http.ListenAndServe(":"+c.AppPort, router); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
