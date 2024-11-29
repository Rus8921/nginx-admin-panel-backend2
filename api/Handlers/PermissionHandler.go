package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
)

func CreatePermissionHandler(context *gin.Context) {
	var permission Permission.Permission
	if err := context.ShouldBindJSON(&permission); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := Permission.CreatePermission(configs.Db, &permission); err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "Permission added successfully", "Permission": permission})
}

func DeletePermissionHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := Permission.DeletePermission(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Permission deleted successfully"})
}

func GetPermissionHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	permission, err := Permission.GetPermission(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Permission found", "Permission": permission})
}

func GetPermissionsAllHandler(context *gin.Context) {
	permissions, err := Permission.GetPermissionAll(configs.Db)
	if err != nil {
		context.JSON(422, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Permissions found", "Permissions": permissions})
}
