package webapp

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/internal/gincontroller"
)

// SetupRouter initializes the router with different route groups
// and their respective middleware.
func SetupRouter() *gin.Engine {
	// Logger and Recovery middleware already attached.
	router := gin.Default()

	// Disable warning "You trusted all proxies, this is NOT safe."
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Could not set trusted proxies: %v", err)
	}

	router.Use(
		// List middlware for static files here. For example:
		// gin.Logger(),
		static.Serve("/", MustFs("")),
	)

	apiRoot := router.Group("/api")
	{
		// List here API without authorization.

		// Authorization required.
		g := apiRoot.Group(
			"/",
			AuthMiddleware(),
			// Add middleware here. For example:
			// middleware.Proxy(),
		)
		{
			g.GET("/hello", gincontroller.ApiHello)
		}
	}

	return router
}
