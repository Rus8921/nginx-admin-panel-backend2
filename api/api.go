package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
)

func registrationHandler(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.RegistrateUser(configs.Db, user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginHandler(context *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := context.ShouldBind(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	user, err := models.LoginUser(configs.Db, credentials.Username, credentials.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials", "details": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User logged in", "user": user})
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

func findHandler(context *gin.Context) {
	var request struct {
		Username string `json:"username"`
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"}) // Проверка на валидность входных данных
		return
	}

	user, err := models.GetUserByUsername(configs.Db, request.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User found", "user": user})
}

func deleteHandler(context *gin.Context) {
	var request struct {
		Id uint `json:"id"`
	}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.DeleteUser(configs.Db, request.Id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User deleted", "userId": request.Id})
}

func main() {

	configs.InitDb()
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/users", findHandler)
		user.POST("/login", loginHandler)
		user.PUT("/users")
		user.DELETE("/users", deleteHandler)
		user.POST("/registration", registrationHandler)
	}
	nginxServer := router.Group("/nginx_server")
	{
		nginxServer.GET("/nginx_server")
		nginxServer.POST("/nginx_server")
		nginxServer.PUT("/nginx_server")
		nginxServer.GET("/nginx_server_list")
		nginxServer.GET("/nginx_server_users")
		nginxServer.DELETE("/nginx_server")
	}

	site := router.Group("/site")
	{
		site.GET("/site")
		site.GET("/site_list")
		site.POST("/site")
		site.PUT("/site")
		site.DELETE("/site")
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
