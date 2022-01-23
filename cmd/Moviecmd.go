package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Movieroute(r *gin.Engine) {
	r.GET("/object", api.Movieimforapi)
	r.GET("/celebrity", api.Personapi)
}

func Moviecommentroute(r *gin.Engine) {
	r.GET("/moviecomment", api.Listcommentapi)
	cm := r.Group("/comment", cookie)
	{
		cm.POST("/parent", api.Commentapi)
		cm.POST("/child", api.Chcommentapi)
	}
}
