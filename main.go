package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db,err:=sql.Open("mysql","root:ljh20040816@tcp(127.0.0.1:3306)/test")
	if(err!=nil){
		fmt.Println(err)
		return
	}
	err=db.Ping()
	if(err!=nil){
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")
}
