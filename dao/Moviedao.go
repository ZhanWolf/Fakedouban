package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"net/http"
)

func QueryMovieimfor(id int) *Struct.Movie {

	M := new(Struct.Movie)
	var psid int
	var persons Struct.Actorinmovie
	var score1 float64
	var score2 float64
	err := Db.QueryRow("select pid,moviename,yyear,introduction,ddate,posterurl,length from movie where id = ?;", id).Scan(&M.Pid, &M.Moviename, &M.Year, &M.Introduction, &M.Date, &M.Poster, &M.Length)
	if err != nil {
		fmt.Println("查询movie出错", err)
		return nil
	}

	sqlStr := "select personid from record_direct where pid=?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows.Next() {
		M.Director = make([]Struct.Actorinmovie, 1)
		err := rows.Scan(psid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename,URL from person where id=?;", psid).Scan(persons.Id, persons.Name, persons.URl)
		M.Director = append(M.Director, persons)
	}
	sqlStr2 := "select personid from record_act where pid=?;"
	rows2, err := Db.Query(sqlStr2, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows2.Next() {
		M.Actor = make([]Struct.Actorinmovie, 1)
		err := rows2.Scan(psid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename,URL from person where id=?;", psid).Scan(persons.Id, persons.Name, persons.URl)
		M.Actor = append(M.Actor, persons)
	}
	sqlStr3 := "select personid from record_act where pid=?;"
	rows3, err := Db.Query(sqlStr3, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows3.Next() {
		M.Scriptwriter = make([]Struct.Actorinmovie, 1)
		err := rows3.Scan(psid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename,URL from person where id=?;", psid).Scan(persons.Id, persons.Name, persons.URl)
		M.Scriptwriter = append(M.Scriptwriter, persons)
	}
	M.Director = M.Director[1:]
	M.Actor = M.Actor[1:]
	M.Scriptwriter = M.Scriptwriter[1:]
	err = Db.QueryRow("select AVG(score) from comment where movie_id = ?;", id).Scan(&score1)
	if err != nil {
		M.Score = 0.0
		return M
	}
	err = Db.QueryRow("select AVG(score) from shortcomment where movie_id = ?;", id).Scan(&score2)
	if err != nil {
		M.Score = 0.0
		return M
	}
	M.Score = (score2 + score1)
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
	var mvid int
	var mvs Struct.Movieinactor
	err := Db.QueryRow("select introduction,birthday,Constellations,chinesename,englishname,birthplace,jobs,posterurl from person where id = ?;", id).Scan(&P.Introduction, &P.Birthday, &P.Constellations, &P.Chinesename, &P.Englishname, &P.Birthplace, &P.Jobs, &P.Jobs, &P.Works, &P.Poster)
	if err != nil {
		fmt.Println("查询movie出错", err)
		return nil
	}

	sqlStr := "select personid from record_direct where pid=?;"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	for rows.Next() {
		P.Works = make([]Struct.Movieinactor, 1)
		err := rows.Scan(mvid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename,URL from person where id=?;", mvid).Scan(mvs.Id, mvs.Name, mvs.URl)
		P.Works = append(P.Works, mvs)
	}
	P.Works = P.Works[1:]
	return P
}

func QueryCooperation(id int) []Struct.Coperson {
	var copersonid int
	sqlStr := "select personid from record_all where personid in (select personid from record_all where pid in(select pid from record_all where personid=?)) and  personid in (select personid from record_all group by personid having count(personid)>2);"
	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	Coperson := make([]Struct.Coperson, 1)
	var Coperson2 Struct.Coperson
	for rows.Next() {
		err := rows.Scan(copersonid)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		Db.QueryRow("select id,chinesename,URL from person where id=?;", copersonid).Scan(Coperson2.Id, Coperson2.Name, Coperson2.URL)
		Coperson = append(Coperson, Coperson2)
	}
	Coperson = Coperson[1:]
	rows.Close()
	return Coperson
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

func Querymovie(id int) bool {
	var moviname string
	err := Db.QueryRow("select moviename from movie where id = ?;", id).Scan(&moviname)
	if err != nil {
		fmt.Println("查询错误", err)
		return false
	}
	return true
}
