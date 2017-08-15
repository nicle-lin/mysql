package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Mysql struct {
	*sql.DB
}

func NewMysql(connStr string) *Mysql{
	db, err := sql.Open("mysql",connStr)
	if err != nil{
		panic(err)
	}
	return &Mysql{db}
}

func (db *Mysql)CheckConn()bool{
	for i:=0; i<=2; i++{
		err := db.Ping()
		if err == nil{
			return true
		}
		time.Sleep(100 * time.Millisecond)
	}
	return false
}