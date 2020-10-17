package main

import (
	"MessageBoard/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/",controller.Index)
	r.GET("/newmessage",controller.NewMessage)
	r.POST("/api/add",controller.AddMessage)
	r.POST("/api/del",controller.DelMessage)
	return r
}
