package controller

import (
	"MessageBoard/common"
	"MessageBoard/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func NewMessage(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"newmessage.tmpl",gin.H{})
}

func AddMessage(ctx *gin.Context)  {
	DB := common.GetDB()
	name:=ctx.PostForm("name")
	content:=ctx.PostForm("content")
	password:=ctx.PostForm("password")
	if len(name)==0 {
		name="匿名用户"
	}
	if len(content)==0||len(content)>150 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"msg":"内容为空或字数超过150",
		})
		return
	}
	if len(password) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"msg":"密码不能为空",
		})
		return
	}
	hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":"加密错误",
		})
		return
	}
	newMessage:=model.Message{
		Name:     name,
		Content:  content,
		Password: string(hasePassowrd),
	}
	DB.Create(&newMessage)
	ctx.JSON(200,gin.H{
		"msg":"创建成功！",
	})
}
func DelMessage(ctx *gin.Context)  {
	db:=common.GetDB()
	var col model.Message

	id:=ctx.PostForm("id")
	db.Where("id=?", id).First(&col)
	password:=ctx.PostForm("password")
	if err := bcrypt.CompareHashAndPassword([]byte(col.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	db.Delete(&col)
	ctx.JSON(200,gin.H{
		"msg":"删除成功",
	})
}
