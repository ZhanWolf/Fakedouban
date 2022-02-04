package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Userroute(r *gin.Engine) {
	us := r.Group("/user")
	{
		us.POST("/login", api.Login)
		us.POST("/Singup", api.Singup)
		us.POST("/Reset", api.Reset)
		us.GET("/clock", cookie, api.Clock)
	}
}

func cookie(c *gin.Context) {
	ck, err := c.Cookie("now_user_login")
	if err != nil {
		fmt.Println(err)
		c.JSON(403, "未登录")
		c.Abort()
	} else {
		c.Set("cookie", ck)
		c.Next()
	}
}
