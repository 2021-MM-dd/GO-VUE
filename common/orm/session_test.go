package orm

import (
	"common/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gogo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Error(err)
	}
	session := New(db)
	if rows, err := session.Raw("select * from user").QueryRows(); err != nil {
		log.Error(err)
	} else {
		var id, name, age, gender string
		for rows.Next() {
			rows.Scan(&id, &name, &age, &gender)
			fmt.Println(id, "===", name, "===", age, "===", gender)
		}
	}

	//session.Raw("select * from user where name = (?) limit (?)", "john", "10").Exec()
}
