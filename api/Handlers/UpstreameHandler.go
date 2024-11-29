package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Upstreams"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
)

func AddUpstreameHandler(context *gin.Context) {
	var upstream Upstreams.Upstream
	if err := context.ShouldBindJSON(&upstream); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := Upstreams.CreateUpstream(configs.Db, &upstream); err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "Upstream added successfully", "Upstream": upstream})
}

func GetUpstreamesAllHandler(context *gin.Context) {
	upstreames, err := Upstreams.GetUpstreamsAll(configs.Db)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Upstreams found", "Upstreams": upstreames})
}

func GetUpstreamHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	upstream, err := Upstreams.GetUpstream(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Upstream found", "Upstream": upstream})
}

func DeleteUpstreamHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := Upstreams.DeleteUpstream(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Upstream deleted successfully"})
}

func UpdateUpstreamHandler(context *gin.Context) {
	var credentials struct {
		Id       uint   `json:"id"`
		Parametr string `json:"parametr"`
		Priority uint   `json:"priority"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	upstream, err := Upstreams.UpdateUpstream(configs.Db, credentials.Id, Upstreams.Upstream{Parametr: credentials.Parametr, Priority: credentials.Priority})
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Upstream updated successfully", "Upstream": upstream})
}
