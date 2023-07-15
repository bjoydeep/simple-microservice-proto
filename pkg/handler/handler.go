package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bjoydeep/simple-microservice-proto/pkg/config"
	"github.com/bjoydeep/simple-microservice-proto/pkg/model"
	"github.com/bjoydeep/simple-microservice-proto/pkg/storage"
	"github.com/bjoydeep/simple-microservice-proto/pkg/transport"
	"github.com/gin-gonic/gin"
)

//var user model.User
//TODO: need to think error handling in general here
//https://gorm.io/docs/error_handling.html

func GetUsers(ctx *gin.Context) {

	var users []model.User
	//gorm generates the query instead of us doing it like
	// rows, err := storage.DBPool.Query(ctx, "select * from public.users")
	storage.DB_.Find(&users)

	// do not get it from model. Get it from DB
	//ctx.JSON(http.StatusOK, model.GetUsers())
	ctx.JSON(http.StatusOK, gin.H{"data": users})

}

func GetUser(ctx *gin.Context) {
	var user model.User
	if err := storage.DB_.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})

}

func AddUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})

	} else {
		//u := model.AddUser(user)

		if result := storage.DB_.Create(&user); result.Error != nil {
			//careful - error message leaks implementation details.
			ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		} else {
			println("Added to DB: -----")
			ctx.JSON(http.StatusCreated, user)

			fmt.Println(user)
			jsonBytes, err := json.Marshal(user)
			if err != nil {
				println("Error marshalling JSON data: ", err.Error())
			}

			client := transport.BrokerClient
			println("Trying to publish -----")
			transport.Publish(client, jsonBytes, config.Cfg.BrokerTopic)
			println("Published --- ")
		}

	}

}
