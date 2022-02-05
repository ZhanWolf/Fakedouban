package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
	"message-board/service"
	"net/http"
	"strconv"
)

func Commentapi(c *gin.Context) {
	var cm Struct.Comment

	cm.From_username, _ = c.Cookie("now_user_login")
	cm.From_id, _ = dao.Queryusername(cm.From_username)
	if cm.From_id == 0 {
		fmt.Println("发生错误")
		c.JSON(http.StatusOK, gin.H{
			"状态": "失败",
		})
		return
	}
	cm.Content = c.PostForm("content")
	score := c.PostForm("score")
	cm.Score, _ = strconv.ParseFloat(score, 64)
	Movieid := c.PostForm("movieid")
	Movieid2, _ := strconv.Atoi(Movieid)
	cm.Movieid = Movieid2

	flag := service.Setcomment(cm, c)
	if flag == false {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"状态":          "评论成功",
		"评论者id":       cm.From_id,
		"评论者username": cm.From_username,
		"评论内容":        cm.Content,
	})
}

func Chcommentapi(c *gin.Context) {
	pid := c.PostForm("pid")
	pid2, _ := strconv.Atoi(pid)
	flag := dao.Querycomment(pid2)
	if flag == false {
		c.JSON(http.StatusOK, gin.H{
			"状态":   "失败",
			"可能原因": "未找到父id的评论",
		})
		fmt.Println("未找到父亲id的评论")
		return
	}
	From_username, _ := c.Cookie("now_user_login")
	From_id, _ := dao.Queryusername(From_username)
	if From_id == 0 {
		fmt.Println("发生错误")
		c.JSON(http.StatusOK, gin.H{
			"状态": "失败",
		})
		return
	}
	Content := c.PostForm("content")
	useful := c.PostForm("useful")

	useful2, _ := strconv.Atoi(useful)

	service.Setchildcomment(pid2, From_id, From_username, Content, useful2)
	c.JSON(http.StatusOK, gin.H{
		"状态":          "评论成功",
		"评论的父亲id":     pid,
		"评论者id":       From_id,
		"评论者username": From_username,
		"评论内容":        Content,
		"有用状态":        useful,
	})
}

func Listcommentapi(c *gin.Context) {
	movieid := c.Query("movieid")
	movieid2, _ := strconv.Atoi(movieid)
	c.JSON(http.StatusOK, service.ListFilmcomment(movieid2))
	c.Next()
}

func Shortcommentapi(c *gin.Context) {
	From_username, _ := c.Cookie("now_user_login")
	From_id, _ := dao.Queryusername(From_username)
	if From_id == 0 {
		fmt.Println("发生错误")
		c.JSON(http.StatusOK, gin.H{
			"状态": "失败",
		})
		return
	}
	content := c.PostForm("content")
	lorw := c.PostForm("lorw")
	lorw1, _ := strconv.Atoi(lorw)
	score := c.PostForm("score")
	score1, _ := strconv.ParseFloat(score, 64)
	Movieid := c.PostForm("movieid")
	Movieid2, _ := strconv.Atoi(Movieid)

	service.Setshortcomment(From_username, From_id, content, lorw1, score1, Movieid2, c)
	c.JSON(http.StatusOK, gin.H{
		"状态":          "评论成功",
		"评论者id":       From_id,
		"评论者username": From_username,
		"评论内容":        content,
		"电影id":        Movieid2,
		"评分":          score1,
	})
}

func Listusecommentapi(c *gin.Context) {
	movie_id := c.PostForm("movie_id")
	movie_id2, _ := strconv.Atoi(movie_id)
	cm := service.ListFlimcommentbyuse(movie_id2)
	c.JSON(http.StatusOK, cm)
}

func Listtimecommentapi(c *gin.Context) {
	movie_id := c.PostForm("movie_id")
	movie_id2, _ := strconv.Atoi(movie_id)
	cm := service.ListFlimcommentbytime(movie_id2)
	c.JSON(http.StatusOK, cm)
}

func Listshortcommentapi(c *gin.Context) {
	movie_id := c.PostForm("movie_id")
	movie_id2, _ := strconv.Atoi(movie_id)
	cm := service.ListFlimshortcommentbyuselimit(movie_id2)
	c.JSON(http.StatusOK, cm)
}

func ListshortcommentapiBytime(c *gin.Context) {
	movie_id := c.PostForm("movie_id")
	movie_id2, _ := strconv.Atoi(movie_id)
	cm := service.ListFilmshortcommentbytime(movie_id2)
	c.JSON(http.StatusOK, cm)
}

func ListshortcommentapiByuse(c *gin.Context) {
	movie_id := c.PostForm("movie_id")
	movie_id2, _ := strconv.Atoi(movie_id)
	cm := service.ListFilmshortcommentbyuse(movie_id2)
	c.JSON(http.StatusOK, cm)
}
