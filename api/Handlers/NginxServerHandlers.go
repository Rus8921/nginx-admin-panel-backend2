package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/NginxServer"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
)

func AddNginxServerHandler(context *gin.Context) {
	var nginxServer NginxServer.NginxServer
	if err := context.ShouldBindJSON(&nginxServer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := NginxServer.CreateNginxSeerver(configs.Db, &nginxServer); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "NginxServer added successfully", "NginxServer": nginxServer})
}

func GetNginxServerHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	nginxServer, err := NginxServer.GetNginxServer(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "NginxServer found", "NginxServer": nginxServer})
}

func GetNginxServersAllHandler(context *gin.Context) {
	servers, err := NginxServer.GetNginxServersAll(configs.Db)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "NginxServers found", "NginxServers": servers})
}

func DeleteNginxServerHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := NginxServer.DeleteNginxServer(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "NginxServer deleted", "NginxServer": credentials.Id})
}

func UpdateNginxServerHandler(context *gin.Context) {
	var credentials struct {
		Id     uint   `json:"id"`
		Ip     string `json:"ip"`
		Domain string `json:"domain"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	server, err := NginxServer.UpdateNginxServer(configs.Db, credentials.Id, NginxServer.NginxServer{Ip: credentials.Ip, Domain: credentials.Domain})
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "NginxServer updated", "NginxServer": server})

}

func ActiveNginxServerHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	var server NginxServer.NginxServer
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := NginxServer.ActivateOrUnactivateServer(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	server, err = NginxServer.GetNginxServer(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error geting server data", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "status of server changed", "NginxServer": server})
}
