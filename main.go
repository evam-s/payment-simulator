package main

import (
	"log"
	"payment-simulator/internal/cache"
	"payment-simulator/internal/config"
	"payment-simulator/internal/db"
	"payment-simulator/internal/folderwatcher"
	"payment-simulator/internal/processing"
	"payment-simulator/internal/routing/incoming"
	_ "payment-simulator/internal/uploads"
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

	fw.InitializeFolderWatcher()
	validation.RegisterCustomValidations()
	router := routing.RoutingSetup()
	processing.SetPacs002CB(config.Pacs002CBURL)

	log.Println("Payments App Started on Port:", config.ServicePort)
	router.Run(":" + config.ServicePort)
}
