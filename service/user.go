package service

import (
	"encoding/json"
	"github.com/firmianavan/we-media-ploy/entity"
	"log"
	"net/http"
	"strconv"
)

func init() {
	//Register("/article/{id}", ShowArticle)
	//Register("/article/{id}/edit", EditArticle)
	//Register("/article/{id}/comment", ShowArticle)
	Register("/u/regist", Regist)
	Register("/u/login", Login)
	Register("/u/{user_id}", GetUserInfo)
	Register("/u/update", GetUserInfo)
}

/*
func EditArticle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hell33o" + r.FormValue("ad")))
}*/

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	user := entity.User{}
	err = entity.QueryUnique(&user, "user_id", id)
	if err != nil {
		if err == entity.EmptyResultError {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
		return
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	json.NewDecoder(r.Body).Decode(&user)

	_, err := entity.UpdateById(&user, "user_id", "created")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
func Regist(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	json.NewDecoder(r.Body).Decode(&user)

	_, err := entity.InsertAndGetId(&user, "user_id", "created")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	json.NewDecoder(r.Body).Decode(&user)
	ret := entity.User{}
	err := entity.QueryUnique(&ret, "email", user.Email)
	if err != nil && err != entity.EmptyResultError {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
		json.NewEncoder(w).Encode(err)
	} else if err == nil && user.Passwd == ret.Passwd {
		//登录成功
		w.WriteHeader(http.StatusOK)
		ret.Passwd = ""
		json.NewEncoder(w).Encode(Resp{Status: "success", Msg: "ok", Data: ret})
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Resp{Status: "fail", Msg: "not matched"})
	}
}
