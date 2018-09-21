package sqlbuilder

import (
	"log"
	"unicode/utf8"
)

type SqlBuilder struct {
	sql      string
	hasWhere bool //默认false表示从句没有where
	hasOrder bool //默认false表示从句没有一个order
}

//创建sql insert语句,
func (SqlBuilder *SqlBuilder) Insert(table string, cols []string) *SqlBuilder {
	sql := "INSERT INTO " + table + " ("
	colsCount := len(cols)
	for _, v := range cols {
		sql += v + ","
	}
	rs := []rune(sql)
	sql = string(rs[:(utf8.RuneCountInString(sql) - 1)])
	sql += ") VALUES ("
	for i := 0; i < colsCount; i++ {
		sql += "?,"
	}
	rs = []rune(sql)
	sql = string(rs[:(utf8.RuneCountInString(sql) - 1)])
	sql += ")"
	SqlBuilder.sql = sql
	return SqlBuilder
}

//创建sql delete
func (SqlBuilder *SqlBuilder) Delete(table string) *SqlBuilder {
	sql := "DELETE FROM " + table
	SqlBuilder.sql = sql
	return SqlBuilder
}

//where从句, like 可以接在运算符里
func (SqlBuilder *SqlBuilder) Where(col string, operator string) *SqlBuilder {
	//如果sql语句中已经有where从句了，则使用and拼接, 否则用where拼接
	var sql string
	if SqlBuilder.hasWhere {
		sql = SqlBuilder.sql + " AND " + col + " " + operator + " ?"
	} else {
		sql = SqlBuilder.sql + " WHERE " + col + " " + operator + " ?"
		SqlBuilder.hasWhere = true
	}
	SqlBuilder.sql = sql
	return SqlBuilder
}

//orWhere从句 like 可以接在运算符里
func (SqlBuilder *SqlBuilder) OrWhere(col string, operator string) *SqlBuilder {
	//如果sql语句中已经有where从句了，则使用and拼接, 否则用where拼接
	var sql string
	if SqlBuilder.hasWhere {
		sql = SqlBuilder.sql + " OR " + col + " " + operator + " ?"
	} else {
		log.Fatalln("orWhere error; this is the first where paragraph")
	}
	SqlBuilder.sql = sql
	return SqlBuilder
}

//where从句, in 可以接在运算符里, n表示in的个数
func (SqlBuilder *SqlBuilder) In(col string, n int) *SqlBuilder {
	//如果sql语句中已经有where从句了，则使用and拼接, 否则用where拼接
	var sql string
	quesStr := "?"
	for i := 1; i < n; i++ {
		quesStr += ",?"
	}
	if SqlBuilder.hasWhere {
		sql = SqlBuilder.sql + " AND " + col + " IN (" + quesStr + ")"
	} else {
		sql = SqlBuilder.sql + " WHERE " + col + " IN (" + quesStr + ")"
		SqlBuilder.hasWhere = true
	}
	SqlBuilder.sql = sql
	return SqlBuilder
}

//where从句, orIn 可以接在运算符里
func (SqlBuilder *SqlBuilder) OrIn(col string, n int) *SqlBuilder {
	//如果sql语句中已经有where从句了，则使用and拼接, 否则用where拼接
	var sql string
	quesStr := "?"
	for i := 1; i < n; i++ {
		quesStr += ",?"
	}
	if SqlBuilder.hasWhere {
		sql = SqlBuilder.sql + " Or " + col + " IN (" + quesStr + ")"
	} else {
		sql = SqlBuilder.sql + " WHERE " + col + " IN (" + quesStr + ")"
		SqlBuilder.hasWhere = true
	}
	SqlBuilder.sql = sql
	return SqlBuilder
}

//order
func (SqlBuilder *SqlBuilder) OrderBy(col string, direction string) *SqlBuilder {
	var sql string
	if SqlBuilder.hasOrder {
		sql = SqlBuilder.sql + "," + col + " " + direction
	} else {
		sql = SqlBuilder.sql + " ORDER BY " + col + " " + direction
		SqlBuilder.hasOrder = true
	}
	SqlBuilder.sql = sql
	return SqlBuilder
}

//order by field (id,1,54,3,2,1)
func (SqlBuilder *SqlBuilder) OrderByField(col string, n int) *SqlBuilder {
	var sql string
	quesStr := "?"
	for i := 1; i < n; i++ {
		quesStr += ",?"
	}
	if !SqlBuilder.hasOrder {
		sql = SqlBuilder.sql + " ORDER BY FIELD (" + col + "," + quesStr + ")"
		SqlBuilder.hasOrder = true
	}
	SqlBuilder.sql = sql
	return SqlBuilder
}

//limit, order语句在limit之前
func (SqlBuilder *SqlBuilder) Limit(offset string, size string) *SqlBuilder {
	sql := SqlBuilder.sql + " LIMIT " + offset + "," + size
	SqlBuilder.sql = sql
	return SqlBuilder
}

//SELECT
func (SqlBuilder *SqlBuilder) Select(cols string, table string) *SqlBuilder {
	sql := "SELECT " + cols + " FROM " + table
	SqlBuilder.sql = sql
	return SqlBuilder
}

//UPDATE
func (SqlBuilder *SqlBuilder) Update(table string, cols []string) *SqlBuilder {
	sql := "UPDATE " + table + " SET "
	for _, v := range cols {
		sql += v + "=?,"
	}
	runes := []rune(sql)
	sql = string(runes[:utf8.RuneCountInString(sql)-1])
	SqlBuilder.sql = sql
	return SqlBuilder
}

//用户自定义复杂sql自行拼接
func (SqlBuilder *SqlBuilder) Raw(sql string) *SqlBuilder {
	SqlBuilder.sql = SqlBuilder.sql + " " + sql
	return SqlBuilder
}

//获取sql
func (SqlBuilder *SqlBuilder) GetSql() string {
	return SqlBuilder.sql
}
