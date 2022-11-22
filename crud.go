package star

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type People struct {
	Id    int
	Name  string
	Slary int
}

func Crud() {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		log.Println(err)
	}
	sql1 := "SELECT *FROM people"
	var people1 People
	err = db.QueryRow(sql1).Scan(&people1.Id, &people1.Name, &people1.Slary)
	fmt.Println("people1 is ", people1)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
