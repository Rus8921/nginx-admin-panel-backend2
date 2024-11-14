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
		nginxServer.GET("/nginx_server_get", Handlers.GetNginxServerHandler)
		nginxServer.POST("/nginx_server_add", Handlers.AddNginxServerHandler)
		nginxServer.POST("/nginx_server_activate", Handlers.ActiveNginxServerHandler)
		nginxServer.PUT("/nginx_server_change", Handlers.UpdateNginxServerHandler)
		nginxServer.GET("/nginx_server_list", Handlers.GetNginxServersAllHandler)
		nginxServer.GET("/nginx_server_users")
		nginxServer.DELETE("/nginx_server_del", Handlers.DeleteNginxServerHandler)
	}

	site := router.Group("/site")
	{
		site.GET("/site_get", Handlers.GetSiteHandler)
		//site.GET("/site_list")
		site.POST("/site", Handlers.AddSiteHandler)
		site.PUT("/site", Handlers.UpdateSiteHandler)
		site.DELETE("/site", Handlers.DeleteSiteHandler)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
