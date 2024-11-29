package Handlers

import (
	"github.com/gin-gonic/gin"
	auth "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Auth"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/User"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
)

func RegistrationUserHandler(context *gin.Context) {
	var user User.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	if err := User.RegistrateUser(configs.Db, user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token", "details": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user, "token": token})
}

func LoginUserHandler(context *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := context.ShouldBind(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}
	user, err := User.LoginUser(configs.Db, credentials.Username, credentials.Password)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in", "token": token})
}

// почему-то так не работает, выглядит интереснее, но не работает
//func findHandler(context *gin.Context) {
//	username, err := strconv.Atoi(context.Param("username"))
//
//	if err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	}
//
//	user, err := models.GetUserByUsername(configs.Db, strconv.Itoa(username))
//	if err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	context.JSON(http.StatusCreated, gin.H{"message": "User found", "user": user})
//}

func FindUserHandler(context *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
	}

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()}) // Проверка на валидность входных данных
		return
	}

	user, err := User.GetUserByUsername(configs.Db, credentials.Username)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User found", "user": user})
}

func DeleteUserHandler(context *gin.Context) {
	var credentials struct {
		Id uint `json:"id"`
	}

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
	}

	if err := User.DeleteUser(configs.Db, credentials.Id); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted", "userId": credentials.Id})
}

func UpdateUserHandler(context *gin.Context) {
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

	user, err := User.UpdateUser(configs.Db, credentials.Username, credentials.Password, User.User{Username: credentials.NewUsername, Email: credentials.NewEmail, HashPassword: credentials.NewPassword})
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func GetUserByUsernameHandler(context *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
	}

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
	}

	user, err := User.GetUserByUsername(configs.Db, credentials.Username)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User found", "user": user})
}
