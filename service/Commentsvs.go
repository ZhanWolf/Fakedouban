package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
	"net/http"
)

func Setcomment(cm Struct.Comment, c *gin.Context) bool {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return false
	}
	flag := dao.Querymovie(cm.Movieid)
	if flag == false {
		c.JSON(http.StatusOK, gin.H{
			"状态":   "失败",
			"可能原因": "没有该电影",
		})
		return false
	}
	err = dao.Insertcomment(cm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
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
	cm := dao.QuerycommentwithoutChild(movieid)
	return cm
}

func ListFlimcommentbytime(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm, _ := dao.Queryusermoviecm(movieid)
	return cm
}

func ListFlimcommentbyuse(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm, _ := dao.QueryusermoviecmbyUse(movieid)
	return cm
}

func ListFlimshortcommentbyuselimit(movieid int) []Struct.Shortcomment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QueryshortcommentbyUsebyLimit(movieid)
	return cm
}

func ListFilmshortcommentbytime(movieid int) []Struct.Shortcomment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QueryshortcommentbyTime(movieid)
	return cm
}

func ListFilmshortcommentbyuse(movieid int) []Struct.Shortcomment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QueryshortcommentbyUse(movieid)
	return cm
}

func ListFlimcommentwihtchild(movieid int) []Struct.Comment {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cm := dao.QuerycommentwithoutChild(movieid)
	return cm
}

func Setshortcomment(fromusername string, fromuerid int, content string, lorw int, score float64, movieid int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Insertshortcomment(fromusername, fromuerid, content, lorw, score, movieid)
}
