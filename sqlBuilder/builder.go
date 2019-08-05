package sqlBuilder

import (
	"fmt"
	"log"
	"sort"
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

type Row = map[string]interface{}

const LEFT_JOIN = "LEFT JOIN"
const JOIN = "JOIN"
const RIGHT_JOIN = "RIGHT JOIN"

const SELECT = "SELECT"
const SELECT_DISTINCT = "SELECT DISTINCT"
const UPDATE = "UPDATE"
const INSERT = "INSERT"
const DELETE = "DELETE"
const INSERT_OR_UPDATE = "INSERT_OR_UPDATE"

var defaultLocation = time.Local

func SetDefaultLocation(loc *time.Location) {
	if loc != nil {
		defaultLocation = loc
	}
}

func GetDefaultLocation() *time.Location {
	return defaultLocation
}

func Table(tableName string) *QueryBuilder {
	return newEmptyQuery().Table(tableName)
}

func newEmptyQuery() *QueryBuilder {
	return &QueryBuilder{
		loc: defaultLocation,
	}
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
	if len(args) < 1 {
		return query
	}
	ret, err := InterpolateParams(query, args, q.loc)
	if err != nil {
		log.Println(err)
		return query
	}
	return strings.Trim(ret, " ")
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

func (q *QueryBuilder) Count(field string) *QueryBuilder {
	return q.Select("COUNT(" + field + ") AS `count`")
}

func (q *QueryBuilder) Delete() *QueryBuilder {
	q.queryType = DELETE
	return q
}

func (q *QueryBuilder) Insert(row Row) *QueryBuilder {
	q.queryType = INSERT
	var fields []string
	var values []string
	for k, _ := range row {
		fields = append(fields, k)
	}
	sort.Strings(fields)
	for _, k := range fields {
		values = append(values, q.Format("?", row[k]))
	}
	q.insertRows = 1
	q.insert = fmt.Sprintf("(%s) VALUES (%s)", strings.Join(fields, ", "), strings.Join(values, ", "))
	return q
}

func (q *QueryBuilder) InsertMany(rows []Row) *QueryBuilder {
	if len(rows) < 1 {
		log.Println("rows number must be greater than 0")
		return q
	}
	q.queryType = INSERT
	var fields []string
	row := rows[0]
	for k, _ := range row {
		fields = append(fields, k)
	}
	sort.Strings(fields)
	var lines []string
	for _, row := range rows {
		var values []string
		for _, k := range fields {
			values = append(values, q.Format("?", row[k]))
		}
		lines = append(lines, fmt.Sprintf("(%s)", strings.Join(values, ", ")))
	}
	q.insertRows = len(rows)
	q.insert = fmt.Sprintf("(%s) VALUES %s", strings.Join(fields, ", "), strings.Join(lines, ", "))
	return q
}

func (q *QueryBuilder) Table(tableName string) *QueryBuilder {
	q.tableName = tableName
	q.tableNameEscaped = tableName
	return q
}

func (q *QueryBuilder) setTableAlias(tableName string, aliasName string) {
	if q.mapAliasToTable == nil {
		q.mapAliasToTable = make(map[string]string)
	}
	if q.mapTableToAlias == nil {
		q.mapTableToAlias = make(map[string]string)
	}
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

func (q *QueryBuilder) Join(tableName string, fields ...string) *QueryBuilder {
	return q.addJoinTable(tableName, JOIN, fields, "")
}

func (q *QueryBuilder) LeftJoin(tableName string, fields ...string) *QueryBuilder {
	return q.addJoinTable(tableName, LEFT_JOIN, fields, "")
}

func (q *QueryBuilder) RightJoin(tableName string, fields ...string) *QueryBuilder {
	return q.addJoinTable(tableName, RIGHT_JOIN, fields, "")
}

func (q *QueryBuilder) On(condition string, args ...Value) *QueryBuilder {
	last := q.joinTables[len(q.joinTables)-1]
	last.on = q.Format(condition, args...)
	q.joinTables[len(q.joinTables)-1] = last
	return q
}

func (q *QueryBuilder) Where(query string, args ...Value) *QueryBuilder {
	return q.And(query, args...)
}

func (q *QueryBuilder) And(query string, args ...Value) *QueryBuilder {
	q.conditions = append(q.conditions, q.Format(query, args...))
	return q
}

func (q *QueryBuilder) Set(update string, args ...Value) *QueryBuilder {
	q.update = append(q.update, q.Format(update, args...))
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
	var sql string
	switch q.queryType {
	case SELECT:
		sql = q.buildSelect(where)
	case SELECT_DISTINCT:
		sql = q.buildSelect(where)
	case UPDATE:
		tail := sqlTailString(where, q.orderBy, q.limit)
		sql = fmt.Sprintf("UPDATE %s SET %s %s", q.tableNameEscaped, strings.Join(q.update, ", "), tail)
	case INSERT:
		sql = fmt.Sprintf("INSERT INTO %s %s", q.tableNameEscaped, q.insert)
	case INSERT_OR_UPDATE:
		sql = fmt.Sprintf("INSERT INTO %s %s ON DUPLICATE KEY UPDATE %s", q.tableNameEscaped, q.insert, strings.Join(q.update, ", "))
	case DELETE:
		tail := sqlTailString(where, q.orderBy, q.limit)
		sql = fmt.Sprintf("DELETE FROM %s %s", q.tableNameEscaped, tail)
	default:
		sql = ""
	}
	return strings.Trim(sql, " ")
}

func (q *QueryBuilder) buildSelect(where string) string {
	var join []string
	if len(q.joinTables) > 0 {
		for _, item := range q.joinTables {
			str := item.joinType + " " + item.table
			a, ok := q.mapTableToAlias[item.table]
			if ok {
				str += " AS " + a
			} else {
				a = item.table
			}
			if item.on != "" {
				str += " ON " + item.on
			}
			if len(item.fields) > 0 {
				q.fields = append(q.fields, item.fields...)
			}
			join = append(join, str)
		}
	}
	if len(q.fields) < 1 {
		q.fields = append(q.fields, "*")
	}
	tail := sqlTailString(strings.Join(join, " "), where, q.orderBy, q.limit)
	table := q.tableNameEscaped
	if q.mapTableToAlias != nil && len(q.mapTableToAlias[q.tableName]) > 0 {
		table += " AS " + q.mapTableToAlias[q.tableName]
	}
	return fmt.Sprintf("%s %s FROM %s %s", q.queryType, strings.Join(q.fields, ", "), table, tail)
}

func Custom(query string, args ...Value) string {
	if len(args) < 1 {
		return strings.Trim(query, " ")
	}
	ret, err := InterpolateParams(query, args, defaultLocation)
	if err != nil {
		log.Println(err)
		return query
	}
	return strings.Trim(ret, " ")
}
