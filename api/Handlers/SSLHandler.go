package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/SSLcertificat"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
)

func GetSSLCertificateHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	sslCertificate, err := SSLcertificat.GetSSL(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SSL certificate found", "SSL certificate": sslCertificate})
}

func GetSSLCertificatesAllHandler(context *gin.Context) {
	sslCertificates, err := SSLcertificat.GetSSLCertificatesAll(configs.Db)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "All SSL certificates found", "SSL certificates": sslCertificates})
}

func AddSSLCertificateHandler(context *gin.Context) {
	var ssl SSLcertificat.SSL
	if err := context.ShouldBindJSON(&ssl); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := SSLcertificat.CreateSSL(configs.Db, &ssl); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "SSL certificate added successfully", "SSL certificate": ssl})
}

func DeletSSLHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := SSLcertificat.DeleteSSL(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SSL certificate deleted"})
}

func UpdateSSLHandler(context *gin.Context) {
	var credentials struct {
		Id      uint   `json:"id"`
		FileCrt string `json:"new-fileCrt"`
		FileKey string `json:"new-fileKey"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	ssl, err := SSLcertificat.UpdateSSL(configs.Db, credentials.Id, SSLcertificat.SSL{FileCrt: credentials.FileCrt, FileKey: credentials.FileKey})
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SSL certificate updated successfully", "SSL certificate": ssl})
}

func ActivateOrUnactivateSSLHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := SSLcertificat.ActivateOrUnactivateSSL(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SSL certificate activated"})
}
