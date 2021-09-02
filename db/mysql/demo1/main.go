package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	driver := "root:admin@tcp(localhost:3306)/330?charset=utf8"

	db, err := sql.Open("mysql", driver)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from user_todo")
	if err != nil {
		log.Fatalln(err)
	}
	rows.Next()
	fmt.Println(rows.Columns())
	fmt.Println(rows.ColumnTypes())
	a := 0
	err = rows.Scan(&a)
	if err != nil {
		log.Fatalln(err)
	}


}

