package main

import (
	"fmt"
	// "os"
	// "path/filepath"
	"payment-simulator/internal/config"
	"payment-simulator/internal/db"
	"payment-simulator/internal/routing"
)

func main() {
	config := config.LoadConfig()
	fmt.Println("Starting App in ", config.ServiceMode, " Mode...")
	db.ConnectMongo(config.DBTECH + "://" + config.DBURL + ":" + config.DBPORT)
	router := routing.RoutingSetup()

	fmt.Println("Payments App Started on Port: ", config.ServicePort)
	router.Run(":" + config.ServicePort)
}

// func ensureDir(path string) error {
// 	dir := filepath.Dir(path)
// 	return os.MkdirAll(dir, os.ModePerm)
// }
