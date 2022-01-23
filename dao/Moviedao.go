package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"net/http"
)

func QueryMovieimfor(id int) *Struct.Movie {

	M := new(Struct.Movie)

	err := Db.QueryRow("select pid,moviename,score,yyear,introduction,ddate,poster from movie where id = ?;", id).Scan(&M.Pid, &M.Moviename, &M.Score, &M.Year, &M.Introduction, &M.Date, &M.Poster)
	if err != nil {
		fmt.Println("查询movie出错", err)
		return nil
	}

	var personid int
	sqlStr := "select personid from record_direct where pid =?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(personid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		M.Director = append(M.Director, personid)
	}

	sqlStr = "select personid from record_act where pid =?;"
	rows, err = Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(personid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		M.Actor = append(M.Actor, personid)
	}

	sqlStr = "select personid from record_script where pid =?;"
	rows, err = Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(personid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		M.Scriptwriter = append(M.Scriptwriter, personid)
	}

	rows.Close()

	return M
}

func QueryMoviepic(id int, c *gin.Context) {
	Mo := new(Struct.Moviepic)
	sqlStr := "select id,pid,url from moviepic where pid=?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&Mo.Id, &Mo.Pid, &Mo.URL)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		func() {
			c.JSON(http.StatusOK, gin.H{
				"Id":    Mo.Id,
				"电影id":  Mo.Pid,
				"图片URL": Mo.URL,
			})
		}()
	}

	rows.Close()

}

func QueryPersonimfor(id int) *Struct.Person {
	P := new(Struct.Person)
	err := Db.QueryRow("select introduction,birthday,Constellations,chinesename,englishname,birthplace,jobs,works,poster from person where id = ?;", id).Scan(&P.Introduction, &P.Birthday, &P.Constellations, &P.Chinesename, &P.Englishname, &P.Birthplace, &P.Jobs, &P.Jobs, &P.Works, &P.Poster)
	if err != nil {
		fmt.Println("查询movie出错", err)
		return nil
	}

	var personid int
	sqlStr := "select pid from record_direct where personid =?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(personid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		P.Works = append(P.Works, personid)
	}
	return P
}

func QueryCooperation(id int) []int {
	copersonidslice := make([]int, 0)
	var copersonid int
	sqlStr := "select personid from record_all where personid in (select personid from record_all where pid in(select pid from record_all where id=?)) and  personid in (select personid from record_all group by personid having count(personid)>2);"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(copersonid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		copersonidslice = append(copersonidslice, copersonid)

	}

	rows.Close()
	return copersonidslice
}

func QueryMoviepic2(id int) []string {
	Moviepiclice := make([]string, 0)
	var moviepicurl string
	sqlStr := "select id,pid,url from moviepic where pid =?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(moviepicurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Moviepiclice = append(Moviepiclice, moviepicurl)

	}

	rows.Close()

	return Moviepiclice
}

func QueryPersonpic(id int) []string {
	Personpicslice := make([]string, 0)
	var personpicurl string
	sqlStr := "select id,pid,url from personpic where pid =?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}

	for rows.Next() {
		err := rows.Scan(personpicurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Personpicslice = append(Personpicslice, personpicurl)

	}

	rows.Close()
	return Personpicslice
}
