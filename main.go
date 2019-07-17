package main

import (
	"chat/ctrl"
	"html/template"
	"log"
	"net/http"
)

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
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	RegisterView()
	http.ListenAndServe(":8080", nil)
}
