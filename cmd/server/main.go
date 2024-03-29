package main

import (
	"asg-2/config"
	"asg-2/order"
	"asg-2/storage/sql"
	"asg-2/transport/rest"
	"log"
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
	router.Logger.Fatal(router.Start(":" + c.AppPort))
}
