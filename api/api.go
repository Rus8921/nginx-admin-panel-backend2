package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/Handlers"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
)

func main() {

	configs.InitDb()
	router := gin.Default()
	//router.Use(
	//	static.Serve("/", webapp.MustFs("")),
	//)
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

	auth := router.Group("")
	auth.Use(Handlers.JWTAuthMiddleware())
	{
		auth.POST("/site/site_activate", Handlers.ActivateOrUnactivateSiteHandler)
	}

	perm := router.Group("")
	perm.Use(Handlers.SetSiteIDMiddleware())
	perm.Use(Handlers.PermissionMiddleware())
	{
		perm.GET("/site_get", Handlers.GetSiteHandler)
	}

	site := router.Group("/site")
	{
		//site.GET("/site_get", Handlers.GetSiteHandler)
		site.GET("/site_list", Handlers.GetSitesAllHandler)
		site.GET("/ssl_list", Handlers.GetAllSSLCertificatesHandler)
		site.POST("/site_add", Handlers.AddSiteHandler)
		//site.POST("/site_activate", Handlers.ActivateOrUnactivateSiteHandler)
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

	permission := router.Group("/permission")
	{
		permission.GET("/permission_get", Handlers.GetPermissionHandler)
		permission.GET("/permission_list", Handlers.GetPermissionsAllHandler)
		permission.POST("/permission_add", Handlers.CreatePermissionHandler)
		permission.DELETE("/permission_del", Handlers.DeletePermissionHandler)
	}

	upstream := router.Group("/upstream")
	{
		upstream.GET("/upstream_get", Handlers.GetUpstreamHandler)
		upstream.GET("/upstream_list", Handlers.GetUpstreamesAllHandler)
		upstream.POST("/upstream_add", Handlers.AddUpstreameHandler)
		upstream.DELETE("/upstream_del", Handlers.DeleteUpstreamHandler)
		upstream.PUT("/upstream_change", Handlers.UpdateUpstreamHandler)
	}

	configuration := router.Group("/configuration")
	{
		configuration.GET("/configuration_get", Handlers.GetConfigurationHandler)
		configuration.GET("/configuration_list", Handlers.GetConfigurationsAllHandler)
		configuration.POST("/configuration_add", Handlers.CreateConfigurationsHandler)
		configuration.DELETE("/configuration_del", Handlers.DeleteConfigurationHandler)
		configuration.PUT("/configuration_change", Handlers.UpdateConfigurationHandler)
	}

	location := router.Group("/location")
	{
		location.GET("/location_get", Handlers.GetLocationHandler)
		location.GET("/location_list", Handlers.GetLocationsAllHandler)
		location.POST("/location_add", Handlers.AddLocarionHandler)
		location.DELETE("/location_del", Handlers.DeleteLocationHandler)
		location.PUT("/location_change", Handlers.UpdateLocationHandler)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
