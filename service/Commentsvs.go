package service

import (
	"fmt"
	"message-board/Struct"
	"message-board/dao"
)

func Setcomment(cm Struct.Comment) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
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
