package ctrl

import (
	"chat/model"
	"chat/service"
	"chat/util"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	mobile := request.PostFormValue("mobile")
	password := request.PostFormValue("password")

	ok := false
	if mobile == "13423567890" && password == "123456" {
		ok = true
	}

	if ok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		util.Success(writer, 0, "")
	} else {
		util.Fail(writer, "authorize failed")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	mobile := request.PostFormValue("mobile")
	password := request.PostFormValue("password")
	nickname := "joker"
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := service.Register(mobile, password, nickname, avatar, sex)
	if err != nil {
		util.Fail(writer, err.Error())
	} else {
		util.Success(writer, user, "")
	}
}
