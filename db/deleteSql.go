package db

import (
	"gonet/base"
	"strings"
)

func deleteSqlStr(sqlData *SqlData) string {
	key := sqlData.Key
	index := strings.LastIndex(key, ",")
	if index != -1 {
		key = key[:index]
	}
	key = strings.Replace(key, ",", " and ", -1)
	return "delete from " + sqlData.Table + " where " + key
}

//--- struct to sql
func DeleteSql(obj interface{}, params ...OpOption) string {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	op := &Op{sqlType: SQLTYPE_DELETE}
	op.applyOpts(params)
	sqlData := &SqlData{}
	getTableName(obj, sqlData)
	parseStructSql(obj, sqlData, op)
	return deleteSqlStr(sqlData)
}
