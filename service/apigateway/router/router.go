package router

import (
	"net/http"
	"netdisk/middleware"
	"netdisk/service/apigateway/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// 静态资源处理
	router.StaticFS("/static/", http.Dir("./static"))
	router.LoadHTMLGlob("./static/view/*")

	user := router.Group("/user")
	{
		user.GET("/signup", handler.SignUpHandler)
		user.POST("/signup", handler.DoSignUpHandler)
		user.GET("/signin", handler.SignInHandler)
		user.POST("/signin", handler.DoSignInHandler)

		user.Use(middleware.AuthHandler())
		user.POST("/info", handler.UserInfoHandler)
	}
	return router
}
