package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/Handlers"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
)

func main() {

	configs.InitDb()
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/user_get", Handlers.FindUserHandler)
		user.POST("/login", Handlers.LoginUserHandler)
		user.PUT("/user_change", Handlers.UpdateUserHandler)
		user.DELETE("/user_del", Handlers.DeleteUserHandler)
		user.POST("/registration", Handlers.RegistrationUserHandler)
	}
	nginxServer := router.Group("/nginx_server")
	{
		nginxServer.GET("/nginx_server", Handlers.GetNginxServerHandler)
		nginxServer.POST("/nginx_server", Handlers.AddNginxServerHandler)
		nginxServer.PUT("/nginx_server", Handlers.UpdateNginxServerHandler)
		nginxServer.GET("/nginx_server_list", Handlers.GetNginxServersAllHandler)
		nginxServer.GET("/nginx_server_users")
		nginxServer.DELETE("/nginx_server", Handlers.DeleteNginxServerHandler)
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
