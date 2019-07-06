package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func userLogin(writer http.ResponseWriter, request *http.Request) {
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
		Return(writer, 0, data, "")
	} else {
		Return(writer, -1, nil, "authorize failed")
	}

}

func Return(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(ret)
}

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		page := v.Name()
		http.HandleFunc(page, func(writer http.ResponseWriter, request *http.Request) {
			if err := tpl.ExecuteTemplate(writer, page, nil); err != nil {
				log.Fatal(err.Error())
			}
		})
	}
}

func main() {
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/user/login", userLogin)
	RegisterView()
	http.ListenAndServe(":8080", nil)
}
