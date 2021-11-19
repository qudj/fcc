package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func initRouter(r *gin.Engine) {
	// health check
	r.GET("health_check", healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(0, "yes")
	ctx.Status(http.StatusOK)
}