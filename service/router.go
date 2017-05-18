package service

import (
	"log"
	"net/http"
	"strings"
)

type mapFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	Mapping       map[string]mapFunc
	StaticHandler http.Handler
}

var Mux Router = Router{
	Mapping:       make(map[string]mapFunc),
	StaticHandler: http.FileServer(http.Dir("/home/van/archive/workspace/go/src/github.com/firmianavan/we-media-ploy/")),
}

//各模块自行注册路由, 允许含有0个或1个url参数,如果如下含有url参数, 将被解析到parseForm中
//如: /article/add, /article/{articleId},,这几个url均被注册时,优先匹配前者
func Register(url string, f mapFunc) {

	//校验:url变量名不能为空, 不能含有一个以上的变量
	if strings.Contains(url, "{}") {
		log.Fatalf("failed to register  url:%s, invalid url,a param name should be in {}", url)
	}
	if strings.Count(url, "{") > 1 {
		log.Fatalf("failed to register  url:%s, invalid url, contains more than one param in the url", url)
	}

	if url[len(url)-1] == '/' {
		trimed := url[:len(url)-1]
		Mux.Mapping[trimed] = f
	}
	Mux.Mapping[url] = f
	log.Printf("url[%s] registered \n", url)
}

func (m Router) statics(w http.ResponseWriter, r *http.Request) bool {
	if strings.Index(r.URL.Path, "/resource/") == 0 {
		//log.Println(r.URL.Path)
		m.StaticHandler.ServeHTTP(w, r)
		return true
	} else {
		return false
	}
}

func (m Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := ""
	val := ""
	if static := m.statics(w, r); static {
		return
	}
	var mapper mapFunc
	for k, f := range m.Mapping {
		if r.URL.Path == k { //完全匹配,直接返回
			r.ParseForm()
			f(w, r)
			return
		} else if !strings.Contains(k, "{") { //配置的不含参数, 且与url不一致, 不匹配跳过
			continue
		} else { //配置的路径含参数, 判断是否与url匹配
			name, val = Match(r.URL.Path, k)
			if name != "" {
				mapper = f
				break
			}
		}
	}
	if name != "" { //参数匹配
		r.ParseForm()
		r.Form[name] = []string{val}
		log.Println(name + ":" + val)
		mapper(w, r)
	} else {
		log.Println("no mapper found for request url:" + r.URL.Path)
		http.NotFound(w, r)
	}
}

//不匹配name为"". 如果匹配, 返回key,value.即路由中定义的变量名, 实际路径中该变量代表的值 如/u/{id} 和/u/2 返回"id"和"2"
func Match(path, target string) (name string, val string) {
	start := strings.Index(target, "{")
	end := strings.Index(target, "}")
	endStr := target[end+1 : len(target)]
	src := path
	if len(src) <= start {
		return "", ""
	}
	val = src[start : len(src)-len(endStr)]
	name = ""
	//如果参数前段相等, 后段相等且url的参数段不含"/",判定为匹配
	//log.Printf("path:%v  target:%v \n start:%d,end %d ,endStr: %s ,val:%s", path, target, start, end, endStr, val)
	if src[0:start] == target[0:start] && src[len(src)-len(endStr):len(src)] == endStr && !strings.Contains(val, "/") {
		name = target[start+1 : end]
	}
	return name, val
}
