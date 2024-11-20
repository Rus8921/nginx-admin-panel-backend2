package Handlers

import (
	"github.com/gin-gonic/gin"
	models "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"log"
	"net/http"
)

func AddSiteHandler(context *gin.Context) {
	var site models.Site
	if err := context.ShouldBindJSON(&site); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	log.Printf("Received site data in AddSiteHandler: %+v", site)

	if err := models.CreateSite(configs.Db, &site); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}

	// Call GeneratorMiddleware logic here
	generator := models.NewGeneratorModel("Template/main.conf", "Template/some_site.conf", "NginxConfigurators/main.conf", "NginxConfigurators/site.conf")
	mainConfPath, siteConfPath, err := generator.CreateConfigCopies()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create config copies", "details": err.Error()})
		return
	}

	log.Printf("Config copies created: mainConfPath=%s, siteConfPath=%s", mainConfPath, siteConfPath)

	if err := generator.UpdateSiteConfig(siteConfPath, site.SiteName, site.Domain); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update site config", "details": err.Error()})
		return
	}

	if err := generator.IncludeSiteConfigInMain(mainConfPath, siteConfPath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to include site config in main", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Site added successfully",
		"Site":    site,
		//"mainConfPath": mainConfPath,
		//"siteConfPath": siteConfPath,
	})
}

func SetSiteIDMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		var credentials struct {
			Id uint `json:"siteID"`
		}
		if err := context.ShouldBindJSON(&credentials); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
			context.Abort()
			return
		}
		context.Set("siteID", credentials.Id)
		context.Next()
	}
}

func GetSiteHandler(context *gin.Context) {
	siteIDValue, exists := context.Get("siteID")
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "siteID not found in context"})
		return
	}

	siteID, ok := siteIDValue.(uint)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid siteID"})
		return
	}

	site, err := models.GetSite(configs.Db, siteID)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Site found", "Site": site})
}

//func GetSiteHandler(context *gin.Context) {
//	siteIDValue, exists := context.Get("siteID")
//	if !exists {
//		context.JSON(http.StatusBadRequest, gin.H{"error": "siteID not found in context"})
//		return
//	}
//
//	siteID, ok := siteIDValue.(uint)
//	if !ok {
//		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid siteID"})
//		return
//	}
//	site, err := models.GetSite(configs.Db, credentials.Id)
//	if err != nil {
//		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
//		return
//	}
//	context.Set("siteID", credentials.Id)
//	context.JSON(http.StatusOK, gin.H{"message": "Site found", "Site": site})
//}

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

//func UpdateSiteHandler(context *gin.Context) {
//	var credentials struct {
//		Id       uint   `json:"id"`
//		Domain   string `json:"domain"`
//		SiteName string `json:"site_name"`
//	}
//	if err := context.ShouldBindJSON(&credentials); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
//		return
//	}
//	var site models.Site
//	if err := context.ShouldBindJSON(&site); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
//		return
//	}
//	updatedSite, err := models.UpdateSite(configs.Db, credentials.Id, site)
//	if err != nil {
//		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{"message": "Site updated", "Site": updatedSite})
//}

func UpdateSiteHandler(context *gin.Context) {
	var credentials struct {
		Id       uint   `json:"id"`
		SiteName string `json:"SiteName"`
		Domain   string `json:"domain"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	site, err := models.UpdateSite(configs.Db, credentials.Id, models.Site{SiteName: credentials.SiteName, Domain: credentials.Domain})
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Site updated", "Site": site})

}

func ActivateOrUnactivateSiteHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	var site models.Site
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := models.ActivateOrUnactivateSite(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	site, err = models.GetSite(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "status of site changed", "Site": site.IsActive})
}

func GetSitesAllHandler(context *gin.Context) {
	sites, err := models.GetSitesAll(configs.Db)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Sites found", "Sites": sites})
}

func GetAllSSLCertificatesHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	ssl, err := models.GetAllSSLCertificates(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SSL certificates found", "SSL": ssl})
}
