package Handlers

import (
	"github.com/gin-gonic/gin"
	models "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"log"
	"net/http"
)

func GeneratorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		var site models.Site
		if err := context.ShouldBindJSON(&site); err != nil {
			log.Printf("Error binding JSON: %v", err)
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
			context.Abort()
			return
		}

		log.Printf("Received site data: %+v", site)

		generator := models.NewGeneratorModel("Template/main.conf", "Template/some_site.conf", "NginxConfigurators/main.conf", "NginxConfigurators/site.conf")
		mainConfPath, siteConfPath, err := generator.CreateConfigCopies()
		if err != nil {
			log.Printf("Error creating config copies: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create config copies", "details": err.Error()})
			context.Abort()
			return
		}

		log.Printf("Config copies created: mainConfPath=%s, siteConfPath=%s", mainConfPath, siteConfPath)

		if err := generator.UpdateSiteConfig(siteConfPath, site.SiteName, site.Domain); err != nil {
			log.Printf("Error updating site config: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update site config", "details": err.Error()})
			context.Abort()
			return
		}

		log.Printf("Site config updated: siteConfPath=%s", siteConfPath)

		if err := generator.IncludeSiteConfigInMain(mainConfPath, siteConfPath); err != nil {
			log.Printf("Error including site config in main: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to include site config in main", "details": err.Error()})
			context.Abort()
			return
		}

		log.Printf("Site config included in main: mainConfPath=%s, siteConfPath=%s", mainConfPath, siteConfPath)
		context.Next()
	}
}
