package main

import (
	"web-IDE-back-end/handler"
	"web-IDE-back-end/middleware"

	"github.com/gin-gonic/gin"
)

func collectRoutes(r *gin.Engine) {
	r.POST("/api/judge", handler.Judge)
	user := r.Group("/api/user")
	{
		user.POST("/login", handler.Login)
		user.POST("/register", handler.Register)
		user.POST("/refresh", handler.Refresh)
	}
	file := r.Group("/api/file")
	{
		file.GET("/file", middleware.JWTAuthMiddleware, func(c *gin.Context) { c.String(200, "this a response, u can only see it after logging in") })
		file.GET("/dir")
		file.POST("/save")
	}
	r.GET("/api/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
