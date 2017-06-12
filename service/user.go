package service

import (
	"encoding/gob"
	"encoding/json"
	"github.com/firmianavan/we-media-ploy/entity"
	"log"
	"net/http"
	"strconv"
)

//var loginTmpl *template.Template

func init() {
	//Register("/article/{id}", ShowArticle)
	//Register("/article/{id}/edit", EditArticle)
	//Register("/article/{id}/comment", ShowArticle)
	Register("/u/regist", Regist)
	Register("/u/login", Login)
	Register("/u/{user_id}", GetUserInfo)
	Register("/u/update", GetUserInfo)
	Register("/uexist/{email}", IfExistEmail)

	//为了在session中存储user对象:
	//There may also be cases where you want to store a complex datatype within a session,
	//such as a struct. Sessions are serialised using the encoding/gob package,
	//so it is easy to register new datatypes for storage in sessions:
	gob.Register(&entity.User{})

	//从主页稍作修改获取一个模板
	/*ori, _ := ioutil.ReadFile("resource/article.html")
		t := strings.Replace(string(ori), `<div id="react_body"></div>`,
			`<div  class="container bs-docs-single-col-container">
				<div class="row">
					<h1>{{.Title}}</h1>
					<hr/>
	           		{{.Body}}
	           	</div>
	        </div>`, -1)
		var err error
		loginTmpl, err = template.New("article").Parse(t)
		if err != nil {
			log.Fatal(err)
		}*/

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
	//cookie自动登录
	session, err := store.Get(r, "marjaho")
	if err != nil {
		log.Fatal(err)
	}
	if session.Values["login"] == true {
		json.NewEncoder(w).Encode(Resp{Status: "success", Msg: "ok", Data: session.Values["user"]})
		return
	}
	//log.Println(session.Values)
	user := entity.User{}
	json.NewDecoder(r.Body).Decode(&user)
	ret := entity.User{}
	err = entity.QueryUnique(&ret, "email", user.Email)
	if err != nil && err != entity.EmptyResultError {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
		json.NewEncoder(w).Encode(err)
	} else if err == nil && user.Passwd == ret.Passwd {
		//初次登录成功
		//记录到session

		session.Values["login"] = true
		session.Values["user"] = ret
		err = session.Save(r, w)
		if err != nil {
			log.Fatal(err)
		}

		ret.Passwd = ""
		json.NewEncoder(w).Encode(Resp{Status: "success", Msg: "ok", Data: ret})
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Resp{Status: "fail", Msg: "not matched"})
	}
}

//检验邮箱是否被占用
func IfExistEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	ret := entity.User{}
	err := entity.QueryUnique(&ret, "email", email)

	if err != nil && err != entity.EmptyResultError {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
		json.NewEncoder(w).Encode(err)
	} else if err == nil {
		//邮箱重复
		json.NewEncoder(w).Encode(Resp{Status: "fail", Msg: "email used"})
	} else { //空集
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Resp{Status: "success", Msg: "ok"})
	}
}
