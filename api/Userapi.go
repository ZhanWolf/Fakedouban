package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/dao"
	"message-board/service"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	jwtgo "github.com/dgrijalva/jwt-go"
	myjwt "message-board/jwt"
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

	username1 := service.UserLoginser(username, password)
	if username1 == "" {
		c.JSON(403, gin.H{
			"code":   "403",
			"reason": "密码错误",
		})
		return
	}
	id, err := dao.Queryusername(username)
	token := generateToken(c, id, username)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"Id":       id,
		"username": username,
		"token":    token,
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

	flag := service.UserSingup(username, password, passwordagain, protectionQ, protectionA)
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

}

func QueryprotectionQ(c *gin.Context) {
	username := c.PostForm("username") //用户名

	err := service.Checkuseraliveser(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "没有找到该用户",
		})
		return
	}
	service.PasswordReset(c, username)
}

func Reset(c *gin.Context) {
	username := Getusernamefromtoken(c)
	password := c.PostForm("password")
	passwordagain := c.PostForm("passwordagain")
	protectionA := c.PostForm("protectionA")

	service.PasswordReset2(c, username, password, protectionA, passwordagain)
}

func Clock(c *gin.Context) {
	username := Getusernamefromtoken(c)
	c.JSON(http.StatusOK, gin.H{
		"hello": username,
		"现在时间":  time.Now(),
	})
}

func Userimfor(c *gin.Context) {
	username := Getusernamefromtoken(c)
	U := service.Listuserimfor(username, c)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, U)
}

func Setuserintroduction(c *gin.Context) {
	username := Getusernamefromtoken(c)
	introduction := c.PostForm("introduction")
	err := service.Setintroduction(username, introduction)
	if err != nil {
		c.JSON(403, gin.H{
			"code":   403,
			"reason": "修改信息失败",
		})
		return
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

func OtherUserimfor(c *gin.Context) {
	id := c.PostForm("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	username, err := dao.Queryuserid(id2)
	if err != nil {
		fmt.Println(err)
	}
	U := service.Listuserimfor(username, c)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
	c.JSON(http.StatusOK, U)
}

func generateToken(c *gin.Context, Id int, Username string) string {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	Id2 := strconv.Itoa(Id)
	claims := myjwt.CustomClaims{
		Id2,
		Username,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(500, gin.H{
			"code":   500,
			"status": -1,
			"msg":    err.Error(),
		})
		return ""
	}

	log.Println(token)

	return token
}

func Getusernamefromtoken(c *gin.Context) string {
	token := c.Request.Header.Get("token")

	log.Print("get token: ", token)

	j := myjwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, _ := j.ParseToken(token)
	return claims.Username
}
