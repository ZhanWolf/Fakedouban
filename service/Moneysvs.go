package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
)

func Setmoneytoother(fromuser, touser, expl string, money int, c *gin.Context) bool {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	mymoney := dao.Querymymoney(fromuser)
	if mymoney < money {
		c.JSON(200, "您的余额不足")
		return true
	}
	money2 := mymoney - money
	dao.Insertothermoney(fromuser, touser, expl, money)
	dao.Updatemoney(money, touser)
	dao.Updatemoney(money2, fromuser)
	return false

}

func Setmoneymyother(fromuser, expl string, money int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	mymoney := dao.Querymymoney(fromuser)
	money2 := mymoney + money
	dao.Insertmymoney(fromuser, expl, money)
	dao.Updatemoney(money2, fromuser)
}

func Checkmoneylist(expl string, c *gin.Context) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	dao.Querylikelist(expl, c)

}
