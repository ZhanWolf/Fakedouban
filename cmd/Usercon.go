package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/api"
	"message-board/jwt"
)

func Userroute(r *gin.Engine) {
	us := r.Group("/user")
	{
		us.POST("/login", api.Login)
		us.POST("/Singup", api.Singup)
		us.POST("/Reset", api.Reset)
		us.POST("/QueryProtectionQ", api.QueryprotectionQ)
		us.GET("/clock", cookie, api.Clock)
		us.GET("/imfor", jwt.JWTAuth(), api.Userimfor)
		us.POST("/change", jwt.JWTAuth(), api.Setuserintroduction)
	}
}

func cookie(c *gin.Context) {
	ck, err := c.Cookie("now_user_login")
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "未登录",
		})
		c.Abort()
	} else {
		c.Set("cookie", ck)
		c.Next()
	}
}
