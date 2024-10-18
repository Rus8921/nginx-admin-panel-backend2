package gincontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func ApiHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello API World"})
}
