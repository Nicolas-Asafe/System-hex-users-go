package gingonic

import (
	"pc/core/services"
	"pc/plugins/gin-gonic/Routers"
	sqlr "pc/plugins/repositorys/sqliteRepository"
	"github.com/gin-gonic/gin"
)



func InitializeApi() {
	r := gin.Default()
	repo := sqlr.SqliteRepoInitialize(InitializeDB())
	serv := services.UserService{Repo:repo}
	Routers.InitializeRouters(serv,r)
	r.Run(":4000")
}
