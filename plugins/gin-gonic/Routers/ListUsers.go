package Routers

import (
	"pc/core/services"
	"github.com/gin-gonic/gin"
)
func ListUsers(serv services.UserService, ctx * gin.Context){
	users,err:=serv.ListUsers()
	if err != nil{
		ctx.Error(err)
		return
	}
	ctx.JSON(200,gin.H{
		"users":users,
	})
}