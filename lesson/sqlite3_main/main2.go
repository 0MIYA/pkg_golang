package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	//postgreSQLとの接続ドライバをimport
	//コード内では使用しないので_アンダーバーでimport
)
var Db *sql.DB

var err error

type Person struct{ //struct定義
	Name string
	Age int
}

func main(){
	Db, err =sql.Open("postgres","user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil{
		log.Panicln(err)
	}
	defer Db.Close()
}
//("postgres","user=workuser  dbname=miyadb  password=monicamode  sslmode=disable")

//C
//R
//U
//D