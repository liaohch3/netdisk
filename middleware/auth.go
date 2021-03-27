package middleware

import (
	"net/http"
	"netdisk/cache"
	"netdisk/consts"

	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := c.Cookie("session")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无法获取session"})
			return
		}
		userID, err := cache.GetUserIDFromSession(session)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "session非法"})
			return
		}

		c.Set(consts.USER_ID, userID)
		c.Next()
	}
}
