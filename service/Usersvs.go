package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"message-board/dao"
	"net/http"
)

func UserLoginser(username string, password string) *http.Cookie {
	dao.OpenDb()
	turepassword := dao.Queryuserpassword(username)
	if turepassword != password {
		return nil
	}

	cookie := &http.Cookie{
		Name:     "now_user_login",
		Value:    username,
		MaxAge:   300,
		Path:     "/",
		HttpOnly: true,
	}
	return cookie
}

func Checkuseraliveser(username string) error {
	dao.OpenDb()
	_, err := dao.Queryusername(username)
	return err
}

func UserSingup(username string, password string, passwordagain string, protectionQ string, protectionA string) (*http.Cookie, bool) {
	dao.OpenDb()

	if passwordagain != password {
		return nil, true
	}

	err := dao.Insertuser(username, password, protectionQ, protectionA)
	if err != nil {
		fmt.Println("注册错误", err)
	}

	cookie := &http.Cookie{
		Name:     "now_user_login",
		Value:    username,
		MaxAge:   300,
		Path:     "/",
		HttpOnly: true,
	}

	return cookie, false
}

func PasswordReset(c *gin.Context, username string, password string, protectionA string, passwordagain string) {
	dao.OpenDb()
	protectionQ, trueprotectionA := dao.Queryprotection(username)
	if protectionQ == "" {
		c.JSON(http.StatusOK, "该账户未设置密保")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			username:  "你好！",
			"宁的密保问题是": protectionQ,
		})
	}

	if trueprotectionA == protectionA && password == passwordagain {
		dao.Updatepassword(password, username)
		c.JSON(http.StatusOK, "密码修改成功")
	} else if passwordagain != password && trueprotectionA == protectionA {
		c.JSON(http.StatusOK, "两次输入密码不相同")
	} else if trueprotectionA != protectionA {
		c.JSON(http.StatusOK, "密保答案错误")
	}
}

func Listuserimfor(username string, c *gin.Context) Struct.Userimfor {
	dao.OpenDb()
	var U Struct.Userimfor
	var err error
	U.Id, err = dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "无此用户")
		return U
	}
	U.Username = username
	U.Introduction, err = dao.Queryintroducton(U.Id)
	U.Cm = dao.QueryUsercm(U.Id)
	U.Scm = dao.QueryUserscm(U.Id)
	U.Looked = dao.Looked(U.Id)
	U.Wanted = dao.Wanted(U.Id)
	return U
}

func Setintroduction(username string, introduction string) error {
	dao.OpenDb()
	Id, err := dao.Queryusername(username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dao.UpdateIntroduction(introduction, Id)
	return nil

}
