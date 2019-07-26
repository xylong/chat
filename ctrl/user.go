package ctrl

import (
	"chat/model"
	"chat/service"
	"chat/util"
	"net/http"
)

var userService service.UserService

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	mobile := request.PostFormValue("mobile")
	password := request.PostFormValue("passwd")

	user, err := userService.Login(mobile, password)

	if err != nil {
		util.Fail(writer, err.Error())
	} else {
		util.Success(writer, user, "")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	mobile := request.PostFormValue("mobile")
	password := request.PostFormValue("passwd")
	nickname := "joker"
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, password, nickname, avatar, sex)
	if err != nil {
		util.Fail(writer, err.Error())
	} else {
		util.Success(writer, user, "")
	}
}
