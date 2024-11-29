package Handlers

import (
	"github.com/gin-gonic/gin"
	models "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
)

func AddLocarionHandler(context *gin.Context) {
	var location models.Location
	if err := context.ShouldBindJSON(&location); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := models.CreateLocation(configs.Db, &location); err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "Location added successfully", "Location": location})
}

func GetLocationHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	location, err := models.GetLocation(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Location found", "Location": location})
}

func GetLocationsAllHandler(context *gin.Context) {
	locations, err := models.GetLocationALL(configs.Db)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Locations found", "Locations": locations})
}

func DeleteLocationHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := models.DeleteLocation(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Location deleted successfully"})
}

func UpdateLocationHandler(context *gin.Context) {
	var credentials struct {
		Id   uint   `json:"id"`
		body string `json:"body"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	location, err := models.UpdateLocation(configs.Db, credentials.Id, models.Location{Body: credentials.body})
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Location updated successfully", "Location": location})
}
