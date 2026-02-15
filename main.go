package main

import (
	"fmt"
	// "os"
	// "path/filepath"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/config"
	"payment-simulator/internal/db"
	"payment-simulator/internal/routing"
	"payment-simulator/internal/validation"
)

func main() {
	config := config.LoadConfig()
	fmt.Println("Starting App in", config.ServiceMode, "Mode...")
	switch config.DBTECH {
	case "mongodb":
		db.ConnectMongo("mongodb://"+config.DBUSER+":"+config.DBPASS+"@"+config.DBURL+":"+config.DBPORT+"/"+config.DBNAME, config.DBNAME)
	}
	switch config.CacheTECH {
	case "redis":
		cache.ConnectRedis(""+config.CacheURL+":"+config.CachePORT, config.CacheUSER, config.CachePASS)
	}
	validation.RegisterCustomValidations()
	router := routing.RoutingSetup()

	fmt.Println("Payments App Started on Port:", config.ServicePort)
	router.Run(":" + config.ServicePort)
}

// func ensureDir(path string) error {
// 	dir := filepath.Dir(path)
// 	return os.MkdirAll(dir, os.ModePerm)
// }
