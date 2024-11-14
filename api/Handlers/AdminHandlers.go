package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Admin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
)

func GetAdminHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	admin, err := Admin.GetAdminById(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Admin found", "Admin": admin})
}

func DeleteAdminHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	err := Admin.DeleteAdmin(configs.Db, credentials.Id)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Admin deleted"})
}

func UpdateAdminHandler(context *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`

		NewUsername string `json:"new-username"`
		NewEmail    string `json:"new-email"`
		NewPassword string `json:"new-password"`
	}

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
	}

	user, err := Admin.UpdateAdmin(configs.Db, credentials.Username, credentials.Password, Admin.Admin{Username: credentials.NewUsername, Email: credentials.NewEmail, HashPassword: credentials.NewPassword})
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func RegistrationAdminHandler(context *gin.Context) {
	var admin Admin.Admin
	if err := context.ShouldBind(&admin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := Admin.RegistrateAdmin(configs.Db, admin); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": admin})
}

func LoginAdminHandler(context *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := context.ShouldBind(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	user, err := Admin.LoginAdmin(configs.Db, credentials.Username, credentials.Password)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in", "user": user})
}
