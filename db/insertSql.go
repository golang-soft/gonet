package db

import (
	"gonet/base"
	"strings"
)

func insertSqlStr(sqlData *SqlData) string {
	sqlname := sqlData.Name
	sqlvalue := sqlData.Value
	index := strings.LastIndex(sqlname, ",")
	if index != -1 {
		sqlname = sqlname[:index]
	}

	index = strings.LastIndex(sqlvalue, ",")
	if index != -1 {
		sqlvalue = sqlvalue[:index]
	}
	sql := "insert into " + sqlData.Table + " (" + sqlname + ") VALUES (" + sqlvalue + ")"
	base.GLOG.Debugf(sql)
	return sql
}

//--- struct to sql
func InsertSql(obj interface{}, params ...OpOption) string {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	op := &Op{sqlType: SQLTYPE_INSERT}
	op.applyOpts(params)
	sqlData := &SqlData{}
	getTableName(obj, sqlData)
	parseStructSql(obj, sqlData, op)
	return insertSqlStr(sqlData)
}
