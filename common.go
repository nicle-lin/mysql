package mysql

import "fmt"

func insertSql(tableName string,sqlMap map[interface{}]interface{})string{
	field := ""
	values := ""

	for k, v := range sqlMap{
		if field != ""{
			field = field + "," + fmt.Sprintf("%s",k)
			values = values + "," + "'" + fmt.Sprintf("%s",v) + "'"
		}else{
			field =  fmt.Sprintf("%s",k)
			values =  "'" + fmt.Sprintf("%s",v) + "'"
		}
	}

	return "insert IGNORE into " + tableName + "(" + field + ") values(" + values + ")"
}
