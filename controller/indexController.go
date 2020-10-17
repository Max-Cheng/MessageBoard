package controller

import (
	"MessageBoard/common"
	"MessageBoard/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	db:=common.GetDB()
	var allMessage []model.Message
	db.Find(&allMessage)
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"users":allMessage,
	})
}


