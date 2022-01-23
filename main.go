package main

import (
	"github.com/gin-gonic/gin"
	"message-board/cmd"
	"net/http"
)

func main() {
	r := gin.Default()
	cmd.Userroute(r)
	cmd.Moneyroute(r)
	cmd.Movieroute(r)
	cmd.Moviecommentroute(r)
	r.GET("/help", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"/message/sendmsg":   "发送留言",
			"/message/sendcom":   "发送评论",
			"/message/listmsg":   "查看留言",
			"/message/listcom":   "查看评论",
			"/message/delete":    "删除评论",
			"/message/update":    "更改留言或评论",
			"/message/mymsg":     "查看给我的信息",
			"/message/nonamemsg": "匿名留言",
			"/message/nonamecom": "匿名评论",
			"/message/likes":     "点赞",
			"/user/login":        "登录",
			"/user/Singup":       "注册",
			"/user/Reset":        "重置密码",
			"/user/clock":        "钟",
			"/money/other":       "转账",
			"/money/me":          "给自己充值",
			"/money/history":     "模糊查询历史记录",
		})
	})
	r.Run(":6060")
}
