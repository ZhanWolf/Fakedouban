package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Moneyroute(r *gin.Engine) {
	mn := r.Group("/money", cookie)
	{
		mn.POST("/other", api.Sendmoneytoother)
		mn.POST("/me", api.Sendmoneytome)
		mn.POST("/history", api.Checkhistory)
	}

}
