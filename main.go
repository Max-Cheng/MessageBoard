package main

import (
	"MessageBoard/common"
	"MessageBoard/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		r.LoadHTMLGlob("./../templates/*")
	} else {
		r.LoadHTMLGlob("templates/*")
	}
	r.Static("/statics", "./statics")
	r = CollectRoute(r)
	r.Run(":" + conf.Accessport)
}

func init() {
	gin.SetMode(gin.DebugMode)
}
