package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
	"net/http"
)

func Setcomment(cm Struct.Comment, c *gin.Context) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	flag := dao.Querymovie(cm.Id)
	if flag == false {
		c.JSON(http.StatusOK, gin.H{
			"状态":   "失败",
			"可能原因": "没有该电影",
		})
		return
	}
	err = dao.Insertcomment(cm)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Setchildcomment(pid int, from_id int, from_username string, content string, useful int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	flag := dao.Insertchcomment(pid, from_id, from_username, content, useful)
	if flag == false {
		fmt.Println("插入评论出错")
	}
}

func ListFilmcomment(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm, _ := dao.Queryusermoviecm(movieid)
	return cm
}
