package main

import (
	"web-IDE-back-end/handler"

	"github.com/gin-gonic/gin"
)

func collectRoutes(r *gin.Engine) {
	r.POST("/api/judge", handler.Judge)
}
