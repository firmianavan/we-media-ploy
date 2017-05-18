package test

import (
	"fmt"
	"github.com/firmianavan/we-media-ploy/service"
	"reflect"
	"testing"
)

type MatchObj struct {
	Path   string
	Target string
	Name   string
	Val    string
}
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	addr string `json:"_"`
}

func TestStrings(test *testing.T) {
	u := User{Id: 1001, Name: "aaa", addr: "bbb"}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			fmt.Printf("%s %s = %v -tag:%s \n",
				t.Field(i).Name,
				t.Field(i).Type,
				v.Field(i).Interface(),
				t.Field(i).Tag)
		}
	}

}
func TestMatch(t *testing.T) {
	cases := []MatchObj{
		MatchObj{
			Path:   "/b",
			Target: "/{cat}",
			Name:   "cat",
			Val:    "b",
		},
		MatchObj{
			Path:   "/ad/d",
			Target: "/{a}/d",
			Name:   "a",
			Val:    "ad",
		},
	}

	for _, obj := range cases {
		n, v := service.Match(obj.Path, obj.Target)
		if n == obj.Name && v == obj.Val {

		} else {
			t.Errorf("%s:%s expected, %s:%s returned", obj.Name, obj.Val, n, v)
		}
	}
}
