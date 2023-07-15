package main

import (
	"github.com/bjoydeep/simple-microservice-proto/pkg/config"
	"github.com/bjoydeep/simple-microservice-proto/pkg/handler"
	"github.com/bjoydeep/simple-microservice-proto/pkg/model"
	"github.com/bjoydeep/simple-microservice-proto/pkg/storage"
	"github.com/bjoydeep/simple-microservice-proto/pkg/transport"
	"github.com/gin-gonic/gin"
)

func loadConfig() {
	err := config.InitConfig()
	if err != nil {
		println("Error while loading the config: ", err)
	} else {
		println("-------------------------------------")
		println("Configuration Details are : ")
		println("-------------------------------------")
		println("BrokerHost is: ", config.Cfg.BrokerHost)
		println("Topic is: ", config.Cfg.BrokerTopic)
		println("Broker Port is: ", config.Cfg.BrokerPort)

		println("Database host is: ", config.Cfg.DBHost)
		println("Database name is: ", config.Cfg.DBName)
		println("Database port is: ", config.Cfg.DBPort)
		println("Database user is: ", config.Cfg.DBUser)
		println("Database SSL is: ", config.Cfg.DBSSL)
		println("Database TMZ is: ", config.Cfg.DBTmz)
		println("-------------------------------------")
	}
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	//load the application config details before doing anything else
	loadConfig()
	// Initialize MQTT client
	transport.SetupTransport()
	//Initialized the DB connections
	storage.SetupStorage()
	//Sets up gorm
	model.SetupModel()

	go transport.Subscribe(transport.BrokerClient, config.Cfg.BrokerTopic)

	// Define API endpoints - to add versioning
	router.GET("/users", handler.GetUsers)
	router.POST("/users", handler.AddUser)
	router.GET("/user/:id", handler.GetUser)

	// Start the server
	router.Run(":8080")
}
