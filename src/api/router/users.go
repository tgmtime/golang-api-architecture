package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tgmtime/golang-api-architecture/api/handler"
	"github.com/tgmtime/golang-api-architecture/api/middleware"
	"github.com/tgmtime/golang-api-architecture/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)

	router.POST("/send-otp", middleware.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
