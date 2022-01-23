package api

import (
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
	"strconv"
)

func Movieimforapi(c *gin.Context) {
	movieid := c.Query("movieid")

	movieid2, _ := strconv.Atoi(movieid)
	M, P := service.Movieinfor(movieid2)

	c.JSON(http.StatusOK, gin.H{
		"电影id":   M.Id,
		"电影pid":  M.Pid,
		"电影名":    M.Moviename,
		"电影年份":   M.Year,
		"电影发行时间": M.Date,
		"电影海报":   M.Poster,
		"电影评分":   M.Score,
		"导演id":   M.Director,
		"编剧id":   M.Scriptwriter,
		"演员id":   M.Actor,
		"剧照":     P,
	})

}

func Personapi(c *gin.Context) {
	personid := c.Query("person")

	personid2, _ := strconv.Atoi(personid)
	P, C, Co := service.Personinfor(personid2)
	c.JSON(http.StatusOK, gin.H{
		"艺人id":       P.Id,
		"艺人中文名":      P.Chinesename,
		"艺人英文名":      P.Englishname,
		"艺人出日":       P.Birthplace,
		"艺人星座":       P.Constellations,
		"艺人出生地":      P.Birthplace,
		"艺人头像":       P.Poster,
		"艺人职位":       P.Jobs,
		"艺人作品id":     P.Works,
		"艺人照片":       C,
		"合作两次以上艺人id": Co,
	})
}
