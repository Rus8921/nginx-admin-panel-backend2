package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Configuration"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
)

func CreateConfigurationsHandler(context *gin.Context) {
	var configuration Configuration.Configuration
	if err := context.ShouldBindJSON(&configuration); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := Configuration.CreateConfiguration(configs.Db, &configuration); err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "Configuration added successfully", "Configuration": configuration})
}

func GetConfigurationsAllHandler(context *gin.Context) {
	configurations, err := Configuration.GetConfigurationAll(configs.Db)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Configurations found", "Configurations": configurations})
}

func GetConfigurationHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	configuration, err := Configuration.GetConfiguration(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Configuration found", "Configuration": configuration})
}

func DeleteConfigurationHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := Configuration.DeleteConfiguration(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Configuration deleted successfully"})
}

func UpdateConfigurationHandler(context *gin.Context) {
	var credentials struct {
		Id       uint   `json:"id"`
		Parametr string `json:"parametr"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	configuration, err := Configuration.UpdateConfiguration(configs.Db, credentials.Id, Configuration.Configuration{Parametrs: credentials.Parametr})
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Configuration updated successfully", "Configuration": configuration})
}
