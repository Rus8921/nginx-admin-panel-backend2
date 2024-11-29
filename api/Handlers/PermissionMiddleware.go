package Handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Auth"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/User"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/configs"
	"net/http"
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

		siteIDValue, exists := context.Get("siteID")
		if !exists {
			context.JSON(http.StatusBadRequest, gin.H{"error": "siteID not found in context"})
			context.Abort()
			return
		}

		siteID, ok := siteIDValue.(uint)
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid siteID"})
			context.Abort()
			return
		}

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
