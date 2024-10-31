package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models"
	"net/http"
)

func registrationHandler(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.RegistrateUser(Db, user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func main() {

	InitDb()
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/users")
		user.POST("/login", func(context *gin.Context) {
			var user models.User
			if err := context.ShouldBind(&user); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := models.CreateUser(Db, &user); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			context.JSON(http.StatusOK, gin.H{"data": user})
		})
		user.PUT("/users")
		user.DELETE("/users")
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
