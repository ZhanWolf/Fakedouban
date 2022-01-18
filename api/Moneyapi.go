package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
	"strconv"
)

func Sendmoneytoother(c *gin.Context) {
	fromuser, _ := c.Cookie("now_user_login")
	touser := c.PostForm("touser")
	money := c.PostForm("money")
	expl := c.PostForm("expl")
	err := service.Checkuseraliveser(touser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "没有该账户")
		return
	}
	money2, _ := strconv.Atoi(money)
	flag := service.Setmoneytoother(fromuser, touser, expl, money2, c)
	if flag {
		return
	}
	c.JSON(200, "转账成功")
}

func Sendmoneytome(c *gin.Context) {
	fromuser, _ := c.Cookie("now_user_login")
	money := c.PostForm("money")
	expl := c.PostForm("expl")
	money2, _ := strconv.Atoi(money)
	service.Setmoneymyother(fromuser, expl, money2)
	c.JSON(http.StatusOK, "充值成功")
}

func Checkhistory(c *gin.Context) {
	expl := c.PostForm("expl")
	service.Checkmoneylist(expl, c)

}
