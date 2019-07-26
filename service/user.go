package service

import (
	"chat/model"
	"chat/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

// 注册
func (s *UserService) Register(mobile, password, nickname, avatar string, sex int) (user model.User, err error) {
	tmp := model.User{}
	_, err = Db.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("this phone number has been used")
	}
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(password, tmp.Salt)
	tmp.Createat = time.Now()
	_, err = Db.InsertOne(&tmp)
	return user, err
}

func (s *UserService) Login(mobile, password string) (user model.User, err error) {
	Db.Where("mobile=?", mobile).Get(&user)
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	if !util.ValidatePasswd(password, user.Salt, user.Passwd) {
		return user, errors.New("密码不正确")
	}
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	user.Token = token
	Db.ID(user.Id).Cols("token").Update(&user)
	return user, nil
}
