package dao

import (
	"fmt"
	"message-board/Struct"
	"time"
)

func Insertcomment(cm Struct.Comment) error {
	time := time.Now()
	_, err := Db.Exec("insert into comment(from_username, from_id, Content, theday, usenum, unusenum, score,movie_id) values(?,?,?,?,0,0,?,?);", cm.From_username, cm.From_id, cm.Content, time, cm.Score, cm.Movieid)
	return err
}

func Insertchcomment(pid int, from_id int, from_username string, content string, useful int) bool {
	time := time.Now()
	var use int
	var unuse int
	if useful == 1 {
		err := Db.QueryRow("select usenum from  comment where id=?;", pid).Scan(&use)
		use++
		_, err = Db.Exec("update comment set usenum=? where id=?;", use, pid)
		if err != nil {
			fmt.Println(err)
		}
	} else if use == 0 {
		err := Db.QueryRow("select unusenum from  comment where id=?;", pid).Scan(&unuse)
		unuse++
		_, err = Db.Exec("update comment set unusenum=? where id=?;", unuse, pid)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("用户绕过客户端！！！，本次操作不予执行")
		return false
	}

	_, err := Db.Exec("insert into childcomment(pid, from_id, from_username, content, theday, Useful) values(?,?,?,?,?,?);", pid, from_id, from_username, content, time, useful)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func Queryusermoviecm(movieid int) ([]Struct.Comment, error) {
	cm := make([]Struct.Comment, 1)
	var cm2 Struct.Comment
	var chcm2 Struct.Childcomment
	var time1 []uint8
	var time2 []uint8
	i := 0
	sqlStr := "select id,from_username, from_id,Content, theday, usenum, unusenum,score,movie_id from comment where movie_id = ?;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		cm2.Child = make([]Struct.Childcomment, 0)
		err = rows.Scan(&cm2.Id, &cm2.From_username, &cm2.From_id, &cm2.Content, &time1, &cm2.Useful, &cm2.Unuseful, &cm2.Score, &cm2.Movieid)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
			return nil, err
		}
		cm2.Time = utos(time1)
		okk, _ := Db.Query("select id, pid, from_id, from_username, content, theday, Useful from childcomment where pid=?;", cm2.Id)
		for okk.Next() {
			err = okk.Scan(&chcm2.Id, &chcm2.Pid, &chcm2.From_id, &chcm2.From_username, &chcm2.Content, &time2, &chcm2.Useful)
			if err != nil {
				fmt.Println("scan failed, err:%v\n", err)
				return nil, err
			}
			chcm2.Time = utos(time2)
			cm2.Child = append(cm2.Child, chcm2)
		}
		okk.Close()
		cm = append(cm, cm2)
		i++
	}

	rows.Close()
	cm = cm[1:]
	return cm, err

}

func Querycomment(pid int) bool {
	var content string
	err := Db.QueryRow("select Content from comment where id = ?;", pid).Scan(&content)
	if err != nil {
		fmt.Println("查询错误", err)
		return false
	}
	return true
}
