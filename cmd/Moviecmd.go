package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Movieroute(r *gin.Engine) {
	r.GET("/object", api.Movieimforapi, api.Moviepicapi, api.Listcommentapi, api.Listshortcommentapi)
	r.GET("/celebrity", api.Personapi, api.Personpic, api.Coperson)
	r.GET("/recommend", api.RealeasingMovieimforapi, api.HotMovieimforapi)
	r.GET("/newhotlist", api.Newhotlist)
	r.POST("/classhot", api.Classhotlist)
	r.POST("/class", api.Classmovie)
	r.POST("/search", api.Searchmovie)
}

func Moviecommentroute(r *gin.Engine) {
	cm := r.Group("/comment", cookie)
	{
		cm.POST("/parent", api.Commentapi)
		cm.POST("/child", api.Chcommentapi)
		cm.POST("/shortcomment", api.Shortcommentapi)
	}
	r.GET("/shortbytime", api.ListshortcommentapiBytime)
	r.GET("/shortbyuse", api.ListshortcommentapiByuse)
	r.GET("/commentbyuse", api.Listusecommentapi)
	r.GET("/commentbytime", api.Listtimecommentapi)
}
