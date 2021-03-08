package main

import (
	"web-IDE-back-end/handler"

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
		file.GET("/file", handler.GetFile)
		file.GET("/dir", handler.GetDir)
		file.POST("/save", handler.SaveFile)
	}
	r.GET("/api/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
