package Handlers

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Auth"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/User"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
	"strconv"
)

func PermissionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
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
