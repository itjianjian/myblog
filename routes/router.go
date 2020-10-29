package routes

import (
	"myblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouters() {

	gin.SetMode(utils.AppMode)
	//r := gin.New()
	//r.Use(middleware.Log())

}
