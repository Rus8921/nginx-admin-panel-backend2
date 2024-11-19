package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Auth"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/User"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"log"
	"net/http"
	"os"
	"strconv"
)

func PermissionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("Error opening log file:", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			context.Abort()
			return
		}
		defer logFile.Close()

		logger := log.New(logFile, "", log.LstdFlags)

		logger.Println("PermissionMiddleware executed")

		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			context.Abort()
			return
		}

		// Проверка токена и извлечение информации о пользователе
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			context.Abort()
			return
		}

		user, err := User.GetUserByUsername(configs.Db, claims.Username)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			context.Abort()
			return
		}
		userID := user.ID

		siteIDstr := context.Param("siteID")
		siteID, err := strconv.Atoi(siteIDstr)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid site ID"})
			context.Abort()
			return
		}

		logger.Printf("Checking permission for userID: %d, siteID: %d\n", userID, siteID)

		hasPermission, err := Permission.CheckPermission(configs.Db, userID, uint(siteID))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking permissions", "details": err.Error()})
			context.Abort()
			return
		}

		if hasPermission == false {
			context.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			context.Abort()
			return
		}

		context.Next()
	}
}
