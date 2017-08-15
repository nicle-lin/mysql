package mysql

import "errors"

var ErrConnection = errors.New("Connection error")

func (db *Mysql)InsertRow(tableName string, sqlMap map[interface{}]interface{})(int64,error) {
	if !db.CheckConn(){
		return 0,ErrConnection
	}
	sqlStr := insertSql(tableName,sqlMap)

	res, err := db.Exec(sqlStr)
	if err != nil{
		return 0, err
	}
	return res.LastInsertId()
}

func InsertRows() {

}

func QueryRow() {

}

func QueryRows() {

}

func DeleteRow() {

}

func DeleteRows() {

}

func UpdateRow() {

}

func UpdateRows() {

}
