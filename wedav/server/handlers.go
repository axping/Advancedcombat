package server

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func WServe() *gin.Engine {
	r := gin.Default()
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

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
		// //创建用户
		// user.POST("/", CreateUser)
		// //登入操作
		user.POST("/:user_name", Login)
	}

}

// func CreateUser(c *gin.Context) {
// }

func Login(c *gin.Context) {

}
