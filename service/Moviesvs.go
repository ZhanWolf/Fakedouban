package service

import (
	"fmt"
	"message-board/Struct"
	"message-board/dao"
)

func Movieinfor(id int) (*Struct.Movie, []string) {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}

	M := dao.QueryMovieimfor(id)
	P := dao.QueryMoviepic2(id)

	return M, P
}

func Personinfor(id int) (*Struct.Person, []string, []int) {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}

	M := dao.QueryPersonimfor(id)
	P := dao.QueryPersonpic(id)
	C := dao.QueryCooperation(id)
	return M, P, C
}
