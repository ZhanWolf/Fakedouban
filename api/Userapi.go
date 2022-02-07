package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
	"net/http"
	"time"
	"unicode/utf8"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := service.Checkuseraliveser(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"code":   "403",
			"reason": "没有该用户",
		})
		return
	}

	cookie := service.UserLoginser(username, password)
	if cookie == nil {
		c.JSON(403, gin.H{
			"code":   "403",
			"reason": "密码错误",
		})
		return
	}
	id, err := dao.Queryusername(username)
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{
		"code":     "200",
		"Id":       id,
		"username": username,
	})

}

func Singup(c *gin.Context) {
	username := c.PostForm("username")           //用户名
	password := c.PostForm("password")           //密码
	passwordagain := c.PostForm("passwordagain") //重复输入密码
	protectionQ := c.PostForm("protectionQ")     //密保问题
	protectionA := c.PostForm("protectionA")     //密保答案

	err := service.Checkuseraliveser(username)
	if err == nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "该用户名已存在",
		})
		return
	}
	if utf8.RuneCountInString(password) <= 3 {
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "密码小于3位",
		})
		return
	}
	if utf8.RuneCountInString(username) < 3 {
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "用户名小于3位",
		})
		return
	}

	cookie, flag := service.UserSingup(username, password, passwordagain, protectionQ, protectionA)
	if flag {
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "两次输入密码不正确",
		})
		return
	}
	id, err := dao.Queryusername(username)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"Id":       id,
		"username": username,
	})
	http.SetCookie(c.Writer, cookie)
}

func Reset(c *gin.Context) {
	username := c.PostForm("username") //用户名
	password := c.PostForm("password") //密码
	passwordagain := c.PostForm("passwordagain")
	protectionA := c.PostForm("protectionA") //密保答案

	err := service.Checkuseraliveser(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "没有找到该用户",
		})
		return
	}
	service.PasswordReset(c, username, password, protectionA, passwordagain)
}

func Clock(c *gin.Context) {
	username, _ := c.Cookie("now_user_login")
	c.JSON(http.StatusOK, gin.H{
		"hello": username,
		"现在时间":  time.Now(),
	})
}

func Userimfor(c *gin.Context) {
	username, _ := c.Cookie("now_user_login")
	U := service.Listuserimfor(username, c)
	U.Code = 200
	c.JSON(http.StatusOK, U)
}

func Setuserintroduction(c *gin.Context) {
	username, _ := c.Cookie("now_user_login")
	introduction := c.PostForm("introduction")
	err := service.Setintroduction(username, introduction)
	if err != nil {
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "修改信息失败",
		})
	}
	id, err := dao.Queryusername(username)
	c.JSON(http.StatusOK, gin.H{
		"code":         "200",
		"performance":  "修改简介成功",
		"id":           id,
		"username":     username,
		"introduction": introduction,
	})
}
