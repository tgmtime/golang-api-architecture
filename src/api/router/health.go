package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tgmtime/golang-api-architecture/api/handler"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	r.GET("/", handler.Health)
}
