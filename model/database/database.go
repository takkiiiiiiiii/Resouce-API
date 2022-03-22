package database

import(
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init(){
	var err error
	accessData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8","root","falcon2002","localhost","3306","BulletinBoard")
	Db, err = sql.Open("mysql", accessData)
	if err != nil {
		fmt.Println("DataBase Access Error !!")
		panic(err)
	}
}