package main

import (
	"github.com/gin-gonic/gin"
)

func WServe() *gin.Engine {
	r := gin.Default()
	//ping
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//用户api分组
	UserGroup(r)
	return r
}

func UserGroup(r *gin.Engine) {
	user := r.Group("/user")
	{
		//创建用户
		user.POST("/", CreateUser)
		//登入操作
		user.POST("/:user_name", Login)
	}
}

func CreateUser(c *gin.Context) {

}

func Login(c *gin.Context) {

}
