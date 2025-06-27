package Routers

import (
	"pc/core/services"
	"github.com/gin-gonic/gin"
)

func InitializeRouters(serv services.UserService, g *gin.Engine) {
	g.GET("/users", func(ctx *gin.Context) {
		ListUsers(serv, ctx)
	})
	g.POST("/users/new", func(ctx *gin.Context) {
		CreateUser(serv, ctx)
	})
	g.DELETE("/users/:id", func(ctx *gin.Context) {
		RemoveUser(serv, ctx)
	})
}
