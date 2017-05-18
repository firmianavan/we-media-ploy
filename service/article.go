package service

import (
	"encoding/json"
	"github.com/firmianavan/we-media-ploy/entity"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var tmpl *template.Template

func init() {
	//Register("/article/{id}", ShowArticle)
	//Register("/article/{id}/edit", EditArticle)
	//Register("/article/{id}/comment", ShowArticle)
	Register("/a/new", NewArticle)
	Register("/a/{article_id}", ShowArticle)
	ori, _ := ioutil.ReadFile("resource/article.html")
	t := strings.Replace(string(ori), `<div id="react_body"></div>`,
		`<div  class="container bs-docs-single-col-container">
			<div class="row"> 
				<h1>{{.Title}}</h1> 
				<hr/>
           		{{.Body}}
           	</div>
        </div>`, -1)
	var err error
	tmpl, err = template.New("article").Parse(t)
	if err != nil {
		log.Fatal(err)
	}
}

/*
func EditArticle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hell33o" + r.FormValue("ad")))
}*/

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	article := entity.Article{}
	id, err := strconv.ParseInt(r.FormValue("article_id"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	err = entity.QueryUnique(&article, "article_id", id)
	if err != nil {
		if err == entity.EmptyResultError {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
		return
	} else {
		tmpl.Execute(w, article)
	}
}

func NewArticle(w http.ResponseWriter, r *http.Request) {
	article := entity.Article{}
	json.NewDecoder(r.Body).Decode(&article)
	//外键约束 TODO
	article.UserId = 1
	article.CategoryId = 1

	_, err := entity.InsertAndGetId(&article, "article_id", "updated", "draft", "created")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(article)
	}
}
