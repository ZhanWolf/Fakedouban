package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"net/http"
	"time"
)

func Insertothermoney(fromuser, touser, expl string, money int) {
	time := time.Now()
	_, err := Db.Exec("insert into money_list(fromuser,touser,money,time,expl) values(?,?,?,?,?);", fromuser, touser, money, time, expl)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Insertmymoney(fromuser, expl string, money int) error {
	time := time.Now()
	_, err := Db.Exec("insert into money_list(fromuser,touser,money,time,expl) values(?,?,?,?,?);", fromuser, fromuser, money, time, expl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Querylikelist(expl string, c *gin.Context) {
	var Mn Struct.Money
	myname, _ := c.Cookie("now_user_login")
	sqlStr := "select id,touser,money,fromuser,time from money_list where expl like  CONCAT('%',?,'%') ;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, expl)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		c.JSON(http.StatusOK, "无交易")
		goto sign1
	}

	for rows.Next() {
		err := rows.Scan(&Mn.Id, &Mn.Touser, &Mn.Howmuch, &Mn.Fromuser, &Mn.Time)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if myname != Mn.Fromuser || myname != Mn.Touser {
			c.JSON(200, "无权查看此次交易")
			goto sign2
		}
		func() {
			time := utos(Mn.Time)
			c.JSON(http.StatusOK, gin.H{
				"该次交易id": Mn.Id,
				"交易给":    Mn.Touser,
				"来自":     Mn.Fromuser,
				"交易金额":   Mn.Howmuch,
				"时间":     time,
			})
		}()
	sign2:
	}

	rows.Close()
sign1:
}

func Updatemoney(money int, username string) error {
	_, err := Db.Exec("update user set money=? where username=?;", money, username)
	return err
}

func Querymymoney(username string) int {
	var U Struct.User
	err := Db.QueryRow("select money from user where username=?;", username).Scan(&U.Money)
	if err != nil {
		fmt.Println("查询错误", err)

	}
	return U.Money
}
