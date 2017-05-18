package test

import (
	"database/sql"
	"fmt"
	"github.com/firmianavan/we-media-ploy/entity"
	"testing"
	"time"
)

func setString(des *string) string {
	ret := *des
	*des = "set"
	return ret
}

func TestEntity(test *testing.T) {
	db, err := sql.Open("mysql", "van:123456@tcp(127.0.0.1:3306)/wmp?parseTime=true&loc=UTC")
	if err != nil {
		test.Fatal(err)
	}
	defer db.Close()
	t := time.Now()
	t = t.AddDate(1, 1, 1)
	/*_, e := entity.DeleteById(db, "user", "'1'", "1")
	if e != nil {
		test.Fatal(e)
	}*/
	/*user := entity.User{Account: "asdf2222", Name: "van2", Created: t}
	id, e := entity.InsertAndGetId(db, &user, "user_id")
	if e != nil {
		test.Fatal(e)
	}*/
	//测试查改, 注意空串和时间, 以及测试时间locale

	u2 := entity.User{}
	err = entity.QueryUnique(db, &u2, "user_id", 9)
	if err != nil {
		test.Fatal(err)
	}
	fmt.Println(u2.Name, u2.Email, u2.Created)
	u2.Email = "163.com"
	affect, err := entity.UpdateById(db, &u2, "user_id")
	if err != nil {
		test.Fatal(err)
	}
	fmt.Printf("%d row affected", affect)
	/*rows, err := db.Query("select name,created from user where  user_id =?", 1)
	if err != nil {
		test.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		name := ""
		created := ""
		rows.Scan(&name, &created)
		test.Fatal(name, created, u2.Created)
	} else {
		test.Fatal("no rows returned")
	}*/

}
