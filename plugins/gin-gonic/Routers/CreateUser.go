package Routers

import (
	"pc/core/models"
	"pc/core/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(serv services.UserService,ctx *gin.Context) {
	var newUser models.User
	err:=ctx.ShouldBindBodyWithJSON(&newUser)
	if err != nil{
		ctx.Error(err)
		return
	}
	serv.NewUser(newUser)
	ctx.JSON(200,gin.H{
		"message":"User created succesfully",
	})
}