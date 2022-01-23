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
	Id            int
	From_id       int
	From_username string
	Content       string
	Score         float64
	Time          string
	Useful        int
	Unuseful      int
	Movieid       int
	Child         []Childcomment
}

type Childcomment struct {
	Id            int
	Pid           int
	From_id       int
	From_username string
	Content       string
	Time          string
	Useful        int
}
