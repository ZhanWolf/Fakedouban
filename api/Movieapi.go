package api

import (
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
	"strconv"
)

func Movieimforapi(c *gin.Context) {
	movieid := c.Query("movie_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Movieinfor(movieid2)

	c.JSON(http.StatusOK, M)

}

func Personapi(c *gin.Context) {
	personid := c.Query("person_id")

	personid2, _ := strconv.Atoi(personid)
	P := service.Personinfor(personid2)
	c.JSON(http.StatusOK, P)
}

func Moviepicapi(c *gin.Context) {
	movieid := c.Query("movie_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Moviepicsvs(movieid2)
	c.JSON(http.StatusOK, M)
}

func Personpic(c *gin.Context) {
	movieid := c.Query("person_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Personpicsvs(movieid2)

	c.JSON(http.StatusOK, M)
}

func Coperson(c *gin.Context) {
	movieid := c.Query("person_id")

	movieid2, _ := strconv.Atoi(movieid)
	M := service.Copersonsvs(movieid2)

	c.JSON(http.StatusOK, M)
}
