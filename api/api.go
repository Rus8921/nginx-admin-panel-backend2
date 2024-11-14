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
		nginxServer.GET("/site_list", Handlers.GetAllSitesOfServerHandler)
		nginxServer.DELETE("/nginx_server_del", Handlers.DeleteNginxServerHandler)
	}

	site := router.Group("/site")
	{
		site.GET("/site_get", Handlers.GetSiteHandler)
		site.GET("/site_list", Handlers.GetSitesAllHandler)
		site.GET("/ssl_list", Handlers.GetAllSSLCertificatesHandler)
		site.POST("/site_add", Handlers.AddSiteHandler)
		site.POST("/site_activate", Handlers.ActivateOrUnactivateSiteHandler)
		site.PUT("/site_change", Handlers.UpdateSiteHandler)
		site.DELETE("/site_del", Handlers.DeleteSiteHandler)

	}

	ssl := router.Group("/ssl")
	{
		ssl.GET("/ssl_get", Handlers.GetSSLCertificateHandler)
		ssl.GET("/ssl_list", Handlers.GetSSLCertificatesAllHandler)
		ssl.POST("/ssl_add", Handlers.AddSSLCertificateHandler)
		ssl.POST("/ssl_activate", Handlers.ActivateOrUnactivateSSLHandler)
		ssl.DELETE("/ssl_del", Handlers.DeletSSLHandler)
		ssl.PUT("/ssl_change", Handlers.UpdateSSLHandler)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/admin_get", Handlers.GetAdminHandler)
		admin.POST("/admin_add", Handlers.RegistrationAdminHandler)
		admin.POST("/admin_login", Handlers.LoginAdminHandler)
		admin.PUT("/admin_change", Handlers.UpdateAdminHandler)
		admin.DELETE("/admin_del", Handlers.DeleteAdminHandler)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
