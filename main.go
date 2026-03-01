package main

import (
	"log"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/config"
	"payment-simulator/internal/db"
	"payment-simulator/internal/routing/incoming"
	"payment-simulator/internal/validation"
)

func main() {
	config := config.LoadConfig()
	log.Println("Starting App in", config.ServiceMode, "Mode...")
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

	log.Println("Payments App Started on Port:", config.ServicePort)
	router.Run(":" + config.ServicePort)
}
