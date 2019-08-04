package sqlBuilder

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

type QueryBuilder struct {
	loc                  *time.Location
	tableName            string
	tableNameEscaped     string
	fields               []string
	conditions           []string
	queryType            string
	update               []string
	insert               string
	insertRows           int
	delete               string
	sql                  string
	sqlTpl               string
	sqlValues            []Value
	orderFields          string
	orderBy              string
	groupBy              string
	offsetRows           int
	limitRows            int
	limit                string
	mapTableToAlias      map[string]string
	mapAliasToTable      map[string]string
	currentJoinTableName string
	joinTables           []joinTableItem
}

type joinTableItem struct {
	table    string
	fields   []string
	joinType string
	on       string
	alias    string
}

const LEFT_JOIN = "LEFT JOIN"
const JOIN = "JOIN"
const RIGHT_JOIN = "RIGHT JOIN"

const SELECT = "SELECT"
const SELECT_DISTINCT = "SELECT DISTINCT"
const UPDATE = "UPDATE"
const INSERT = "INSERT"
const DELETE = "DELETE"
const CUSTOM = "CUSTOM"
const INSERT_OR_UPDATE = "INSERT_OR_UPDATE"

func Table(tableName string) *QueryBuilder {
	return &QueryBuilder{tableName: tableName}
}

func newEmptyQuery() *QueryBuilder {
	return &QueryBuilder{}
}

func Update() *QueryBuilder {
	return newEmptyQuery().Update()
}

func Select(fields ...string) *QueryBuilder {
	return newEmptyQuery().Select(fields...)
}

func SelectDistinct(fields ...string) *QueryBuilder {
	return newEmptyQuery().SelectDistinct(fields...)
}

func Insert(row interface{}) *QueryBuilder {
	return newEmptyQuery().Insert(row)
}

func Delete() *QueryBuilder {
	return newEmptyQuery().Delete()
}

func (q *QueryBuilder) Clone() (*QueryBuilder, error) {
	ret := newEmptyQuery()
	err := copier.Copy(ret, q)
	return ret, err
}

func (q *QueryBuilder) Location(loc *time.Location) *QueryBuilder {
	q.loc = loc
	return q
}

func (q *QueryBuilder) Format(query string, args ...Value) string {
	ret, err := InterpolateParams(query, args, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.Trim(ret, "")
}

func (q *QueryBuilder) Update() *QueryBuilder {
	q.queryType = UPDATE
	return q
}

func (q *QueryBuilder) OnDuplicateKeyUpdate() *QueryBuilder {
	q.queryType = INSERT_OR_UPDATE
	return q
}

func (q *QueryBuilder) Select(fields ...string) *QueryBuilder {
	q.queryType = SELECT
	q.Fields(fields...)
	return q
}

func (q *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	q.fields = append(q.fields, fields...)
	return q
}

func (q *QueryBuilder) SelectDistinct(fields ...string) *QueryBuilder {
	q.queryType = SELECT_DISTINCT
	q.Fields(fields...)
	return q
}

func (q *QueryBuilder) Count(name string, field string) *QueryBuilder {
	q.queryType = SELECT
	q.fields = append(q.fields, "COUNT("+field+") AS "+name)
	return q
}

func (q *QueryBuilder) Delete() *QueryBuilder {
	q.queryType = DELETE
	return q
}

func (q *QueryBuilder) Insert(row interface{}) *QueryBuilder {
	q.queryType = INSERT
	return q
}

func (q *QueryBuilder) Table(tableName string) *QueryBuilder {
	q.tableName = tableName
	q.tableNameEscaped = tableName
	return q
}

func (q *QueryBuilder) Into(tableName string) *QueryBuilder {
	return q.Table(tableName)
}

func (q *QueryBuilder) From(tableName string) *QueryBuilder {
	return q.Table(tableName)
}

func (q *QueryBuilder) setTableAlias(tableName string, aliasName string) {
	q.mapAliasToTable[aliasName] = tableName
	q.mapTableToAlias[tableName] = aliasName
}

func (q *QueryBuilder) addJoinTable(tableName string, joinType string, fields []string, alias string) *QueryBuilder {
	q.currentJoinTableName = tableName
	q.joinTables = append(q.joinTables, joinTableItem{
		table:    tableName,
		fields:   fields,
		joinType: joinType,
		on:       "",
		alias:    alias,
	})
	return q
}

func (q *QueryBuilder) As(name string) *QueryBuilder {
	tableName := q.currentJoinTableName
	if tableName == "" {
		tableName = q.tableName
	}
	q.setTableAlias(tableName, name)
	if len(q.joinTables) > 0 {
		q.joinTables[len(q.joinTables)-1].alias = name
	}
	return q
}

func (q *QueryBuilder) Join(tableName string, fields []string) *QueryBuilder {
	return q.addJoinTable(tableName, JOIN, fields, "")
}

func (q *QueryBuilder) LeftJoin(tableName string, fields []string) *QueryBuilder {
	return q.addJoinTable(tableName, LEFT_JOIN, fields, "")
}

func (q *QueryBuilder) RightJoin(tableName string, fields []string) *QueryBuilder {
	return q.addJoinTable(tableName, RIGHT_JOIN, fields, "")
}

func (q *QueryBuilder) On(condition string, args ...Value) *QueryBuilder {
	last := q.joinTables[len(q.joinTables)-1]
	last.on = q.Format(condition, args)
	return q
}

func (q *QueryBuilder) Where(query string, args ...Value) *QueryBuilder {
	return q.And(query, args...)
}

func (q *QueryBuilder) And(query string, args ...Value) *QueryBuilder {
	q.conditions = append(q.conditions, q.Format(query, args))
	return q
}

func (q *QueryBuilder) Set(update string, args ...Value) *QueryBuilder {
	q.update = append(q.update, q.Format(update, args...))
	return q
}

func (q *QueryBuilder) Custom(query string, args ...Value) *QueryBuilder {
	q.queryType = CUSTOM
	q.sqlTpl = query
	q.sqlValues = args
	return q
}

func (q *QueryBuilder) OrderBy(tpl string, args ...Value) *QueryBuilder {
	q.orderFields = q.Format(tpl, args...)
	q.orderBy = "ORDER BY " + q.orderFields
	return q
}

func (q *QueryBuilder) GroupBy(tpl string, args ...Value) *QueryBuilder {
	q.groupBy = "GROUP BY " + q.orderFields
	return q
}

func (q *QueryBuilder) Having(tpl string, args ...Value) *QueryBuilder {
	q.groupBy += " HAVING " + q.Format(tpl, args...)
	return q
}

func (q *QueryBuilder) Skip(n int) *QueryBuilder {
	q.offsetRows = n
	q.limit = sqlLimitString(q.offsetRows, q.limitRows)
	return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limitRows = n
	q.limit = sqlLimitString(q.offsetRows, q.limitRows)
	return q
}

func (q *QueryBuilder) Build() string {
	where := ""
	if len(q.conditions) > 0 {
		where = "WHERE " + strings.Join(q.conditions, " AND ")
	}
	switch q.queryType {
	case SELECT:
		return q.buildSelect(where)
	case SELECT_DISTINCT:
		return q.buildSelect(where)
	case UPDATE:
		return fmt.Sprintf("INSERT INTO %s %s", q.tableNameEscaped, strings.Join(q.update, ", "))
	case INSERT:
		tail := strings.Join([]string{where, q.orderBy, q.limit}, " ")
		return fmt.Sprintf("UPDATE %s SET %s %s", q.tableNameEscaped, q.insert, tail)
	case INSERT_OR_UPDATE:
		return fmt.Sprintf("INSERT INTO %s %s ON DUPLICATE KEY UPDATE %s", q.tableNameEscaped, q.insert, strings.Join(q.update, ", "))
	case DELETE:
		tail := strings.Join([]string{where, q.orderBy, q.limit}, " ")
		return fmt.Sprintf("DELETE FROM %s %s", q.tableNameEscaped, tail)
	case CUSTOM:
		return q.Format(q.sqlTpl, q.sqlValues...)
	}
}

func (q *QueryBuilder) buildSelect(where string) string {

}
