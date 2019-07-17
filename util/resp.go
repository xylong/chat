package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Fail(w http.ResponseWriter, msg string) {
	Return(w, -1, nil, msg)
}

func Success(w http.ResponseWriter, data interface{}, msg string) {
	Return(w, 0, data, msg)
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
