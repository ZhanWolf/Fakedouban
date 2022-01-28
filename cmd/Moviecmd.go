package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Movieroute(r *gin.Engine) {
	r.GET("/object", api.Movieimforapi, api.Moviepicapi, api.Listcommentapi, api.Listshortcommentapi)
	r.GET("/celebrity", api.Personapi, api.Personpic, api.Coperson)

}

func Moviecommentroute(r *gin.Engine) {
	cm := r.Group("/comment", cookie)
	{
		cm.POST("/parent", api.Commentapi)
		cm.POST("/child", api.Chcommentapi)
		cm.POST("/shortcomment", api.Shortcommentapi)
		cm.GET("/shortbytime", api.ListshortcommentapiBytime)
		cm.GET("/shortbyuse", api.ListshortcommentapiByuse)
		cm.GET("/commentbyuse", api.ListshortcommentapiBytime)
		cm.GET("/commentbytime", api.Listtimecommentapi)
	}
}
