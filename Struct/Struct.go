package Struct

import "time"

type User struct {
	Id          int
	Username    string
	Password    string
	ProtectionQ string
	ProtectionA string
	Money       int
}

type Message struct {
	Id             int
	Tousername     string
	Fromusername   string
	Time           []uint8
	Messagecontent string
}

type Money struct {
	Id       int
	Touser   string
	Fromuser string
	Time     []uint8
	Howmuch  int
}

type Movie struct {
	Id           int
	Pid          int
	Introduction string
	Poster       string
	Year         int
	Date         string
	Moviename    string
	Score        float64
	Actor        []int
	Director     []int
	Scriptwriter []int
}

type Person struct {
	Id             int
	Introduction   string
	Birthday       time.Time
	Constellations string
	Chinesename    string
	Englishname    string
	Birthplace     string
	Jobs           string
	Works          []int
	Poster         string
}

type Record struct {
	Id       int
	Pid      int
	Personid int
}

type Moviepic struct {
	Id  int
	Pid int
	URL string
}

type Comment struct {
	Id            int            `json:"评论id"`
	From_id       int            `json:"评论者id"`
	From_username string         `json:"评论者用户名"`
	Content       string         `json:"评论的内容"`
	Score         float64        `json:"评论的分数"`
	Time          string         `json:"评论的时间"`
	Useful        int            `json:"有用数"`
	Unuseful      int            `json:"无用数"`
	Movieid       int            `json:"电影的id"`
	Child         []Childcomment `json:"子评论"`
}

type Childcomment struct {
	Id            int    `json:"子评论的id"`
	Pid           int    `json:"父评论的id"`
	From_id       int    `json:"评论者的id"`
	From_username string `json:"评论者的用户名"`
	Content       string `json:"评论的内容"`
	Time          string `json:"评论的时间"`
	Useful        int    `json:"感觉是否有用"`
}
