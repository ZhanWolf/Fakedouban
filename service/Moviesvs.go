package service

import (
	"fmt"
	"message-board/Struct"
	"message-board/dao"
)

func Movieinfor(id int) *Struct.Movie {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}

	M := dao.QueryMovieimfor(id)

	return M
}

func Personinfor(id int) *Struct.Person {
	err := dao.OpenDb()

	if err != nil {
		fmt.Println(err)
	}

	M := dao.QueryPersonimfor(id)

	return M
}

func Checkmoviealiveser(id int) bool {
	dao.OpenDb()
	flag := dao.Querymovie(id)
	if flag {
		fmt.Println("未找到")
	}
	return flag
}

func Moviepicsvs(id int) []string {
	dao.OpenDb()
	P := dao.QueryMoviepic2(id)

	return P
}

func Personpicsvs(id int) []string {
	dao.OpenDb()
	P := dao.QueryPersonpic(id)

	return P
}

func Copersonsvs(id int) []Struct.Coperson {
	dao.OpenDb()
	C := dao.QueryCooperation(id)

	return C
}
