package service

import (
	"github.com/firmianavan/we-media-ploy/entity"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

type Resp struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

var store = sessions.NewCookieStore([]byte("djhdo8er-dsakdf-ieoqmxdbnurt"))

//对简单无嵌套结构的填充
func Fullfil(r *http.Request, mapper entity.RowMaper) {
	_, cols := mapper.RowMap()
	r.ParseForm()
	for k, v := range cols {
		if val := r.FormValue(k); val != "" {
			scan(val, v.V)
		}
	}

}

func scan(src string, des interface{}) (err error) {
	switch is := des.(type) {
	case *int:
		*is, err = strconv.Atoi(src)
	case *int32:
		tmp, err := strconv.Atoi(src)
		if err != nil {
			*is = int32(tmp)
		}
	case *int64:
		tmp, err := strconv.Atoi(src)
		if err != nil {
			*is = int64(tmp)
		}
	case *float32:
		tmp, err := strconv.ParseFloat(src, 32)
		if err != nil {
			*is = float32(tmp)
		}
	case *float64:
		*is, err = strconv.ParseFloat(src, 64)
	case *string:
		*is = src
	default:
		//*col.V = src
	}
	return err
}
