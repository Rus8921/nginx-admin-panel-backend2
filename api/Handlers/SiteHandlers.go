package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
)

func AddSiteHandler(context *gin.Context) {
	var site models.Site
	if err := context.ShouldBindJSON(&site); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := models.CreateSite(configs.Db, &site); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Site added successfully", "Site": site})
}

func GetSiteHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	site, err := models.GetSite(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Site found", "Site": site})
}

func DeleteSiteHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := models.DeleteSite(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Site deleted"})
}

func UpdateSiteHandler(context *gin.Context) {
	var credentials struct {
		Id       uint   `json:"id"`
		Domain   string `json:"domain"`
		SiteName string `json:"site_name"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	var site models.Site
	if err := context.ShouldBindJSON(&site); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	updatedSite, err := models.UpdateSite(configs.Db, credentials.Id, site)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Site updated", "Site": updatedSite})
}

func ActivateOrUnactivateSiteHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := models.ActivateOrUnactivateSite(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Site activated/unactivated"})
}
