package service

import (
	"chat/model"
	"chat/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 注册
func Register(mobile, password, nickname, avatar string, sex int) (user model.User, err error) {
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

func Login() {

}
