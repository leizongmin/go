package sqlutil

import (
	"database/sql/driver"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

type QueryBuilder struct {
	quoteIdentifier      func(string) string
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
	quoteLiteral         func([]byte, string) []byte
	returningFields      []string
	conflictKey          string
}

type joinTableItem struct {
	table    string
	fields   []string
	joinType string
	on       string
	alias    string
}

const JOIN = "JOIN"
const LEFT_JOIN = "LEFT JOIN"
const RIGHT_JOIN = "RIGHT JOIN"
const INNER_JOIN = "INNER JOIN"

const SELECT = "SELECT"
const SELECT_DISTINCT = "SELECT DISTINCT"
const UPDATE = "UPDATE"
const INSERT = "INSERT"
const DELETE = "DELETE"
const INSERT_OR_UPDATE_DUPLICATE = "INSERT_OR_UPDATE_DUPLICATE"
const INSERT_OR_UPDATE_CONFLICT = "INSERT_OR_UPDATE_CONFLICT"

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
	return NewEmptyQuery().Table(tableName)
}

func NewEmptyQuery() *QueryBuilder {
	return &QueryBuilder{
		loc: defaultLocation,
	}
}

func (q *QueryBuilder) Init(quoteIdentifier func(id string) string, quoteLiteral func([]byte, string) []byte) *QueryBuilder {
	q.quoteIdentifier = quoteIdentifier
	q.quoteLiteral = quoteLiteral
	return q
}

func (q *QueryBuilder) Clone() *QueryBuilder {
	// mapTableToAlias      map[string]string
	// mapAliasToTable      map[string]string
	ret := &QueryBuilder{
		quoteIdentifier:      q.quoteIdentifier,
		loc:                  q.loc,
		tableName:            q.tableName,
		tableNameEscaped:     q.tableNameEscaped,
		queryType:            q.queryType,
		insert:               q.insert,
		insertRows:           q.insertRows,
		delete:               q.delete,
		orderFields:          q.orderFields,
		orderBy:              q.orderBy,
		groupBy:              q.groupBy,
		offsetRows:           q.offsetRows,
		limitRows:            q.limitRows,
		limit:                q.limit,
		currentJoinTableName: q.currentJoinTableName,
		conflictKey:          q.conflictKey,
	}
	copy(ret.fields, q.fields)
	copy(ret.returningFields, q.returningFields)
	copy(ret.conditions, q.conditions)
	copy(ret.update, q.update)
	copy(ret.joinTables, q.joinTables)
	if q.mapTableToAlias != nil {
		ret.mapTableToAlias = make(map[string]string)
		for k, v := range q.mapTableToAlias {
			ret.mapTableToAlias[k] = v
		}
	}
	if q.mapAliasToTable != nil {
		ret.mapAliasToTable = make(map[string]string)
		for k, v := range q.mapAliasToTable {
			ret.mapAliasToTable[k] = v
		}
	}
	return ret
}

func (q *QueryBuilder) QuoteIdentifier(id string) string {
	if q.quoteIdentifier == nil {
		return id
	}
	if id == "*" {
		return id
	}
	if strings.IndexByte(id, '(') != -1 {
		return id
	}
	return q.quoteIdentifier(id)
}

func (q *QueryBuilder) Location(loc *time.Location) *QueryBuilder {
	q.loc = loc
	return q
}

func (q *QueryBuilder) Format(query string, args ...driver.Value) string {
	if len(args) < 1 {
		return query
	}
	ret, err := InterpolateParams(query, args, q.loc, q.quoteLiteral)
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
	q.queryType = INSERT_OR_UPDATE_DUPLICATE
	return q
}

func (q *QueryBuilder) OnConflictDoUpdate(key string) *QueryBuilder {
	q.queryType = INSERT_OR_UPDATE_CONFLICT
	q.conflictKey = key
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
	return q.Select("COUNT(" + field + ") AS " + q.QuoteIdentifier("count"))
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
	for i, k := range fields {
		fields[i] = q.QuoteIdentifier(k)
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
	for i, k := range fields {
		fields[i] = q.QuoteIdentifier(k)
	}
	q.insertRows = len(rows)
	q.insert = fmt.Sprintf("(%s) VALUES %s", strings.Join(fields, ", "), strings.Join(lines, ", "))
	return q
}

func (q *QueryBuilder) Table(tableName string) *QueryBuilder {
	q.tableName = tableName
	q.tableNameEscaped = q.QuoteIdentifier(tableName)
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

func (q *QueryBuilder) InnerJoin(tableName string, fields ...string) *QueryBuilder {
	return q.addJoinTable(tableName, INNER_JOIN, fields, "")
}

func (q *QueryBuilder) On(condition string, args ...driver.Value) *QueryBuilder {
	last := q.joinTables[len(q.joinTables)-1]
	last.on = q.Format(condition, args...)
	q.joinTables[len(q.joinTables)-1] = last
	return q
}

func (q *QueryBuilder) Where(query string, args ...driver.Value) *QueryBuilder {
	return q.And(query, args...)
}

func (q *QueryBuilder) And(query string, args ...driver.Value) *QueryBuilder {
	q.conditions = append(q.conditions, q.Format(query, args...))
	return q
}

func (q *QueryBuilder) WhereRow(row Row) *QueryBuilder {
	return q.AndRow(row)
}

func (q *QueryBuilder) AndRow(row Row) *QueryBuilder {
	var fields []string
	for k, _ := range row {
		fields = append(fields, k)
	}
	sort.Strings(fields)
	for _, k := range fields {
		q.And(q.QuoteIdentifier(k)+"=?", row[k])
	}
	return q
}

func (q *QueryBuilder) Set(update string, args ...driver.Value) *QueryBuilder {
	q.update = append(q.update, q.Format(update, args...))
	return q
}

func (q *QueryBuilder) SetRow(row Row) *QueryBuilder {
	var fields []string
	for k, _ := range row {
		fields = append(fields, k)
	}
	sort.Strings(fields)
	for _, k := range fields {
		q.Set(q.QuoteIdentifier(k)+"=?", row[k])
	}
	return q
}

func (q *QueryBuilder) OrderBy(tpl string, args ...driver.Value) *QueryBuilder {
	q.orderFields = q.Format(tpl, args...)
	q.orderBy = "ORDER BY " + q.orderFields
	return q
}

func (q *QueryBuilder) GroupBy(tpl string, args ...driver.Value) *QueryBuilder {
	q.groupBy = "GROUP BY " + q.Format(tpl, args...)
	return q
}

func (q *QueryBuilder) Having(tpl string, args ...driver.Value) *QueryBuilder {
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

func (q *QueryBuilder) Returning(fields ...string) *QueryBuilder {
	q.returningFields = append(q.returningFields, fields...)
	return q
}

func (q *QueryBuilder) ReturningAll() *QueryBuilder {
	q.returningFields = append(q.returningFields, "*")
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
		sql = q.buildUpdate(tail)
	case INSERT:
		sql = q.buildInsert(where)
	case INSERT_OR_UPDATE_DUPLICATE:
		sql = fmt.Sprintf("INSERT INTO %s %s ON DUPLICATE KEY UPDATE %s", q.tableNameEscaped, q.insert, strings.Join(q.update, ", "))
	case INSERT_OR_UPDATE_CONFLICT:
		sql = fmt.Sprintf("INSERT INTO %s %s ON CONFLICT(%s) DO UPDATE SET %s", q.tableNameEscaped, q.insert, q.conflictKey, strings.Join(q.update, ", "))
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
			str := item.joinType + " " + q.QuoteIdentifier(item.table)
			a, ok := q.mapTableToAlias[item.table]
			if ok {
				str += " AS " + q.QuoteIdentifier(a)
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
	for i, v := range q.fields {
		q.fields[i] = q.QuoteIdentifier(v)
	}
	tail := sqlTailString(strings.Join(join, " "), where, q.groupBy, q.orderBy, q.limit)
	table := q.tableNameEscaped
	if q.mapTableToAlias != nil && len(q.mapTableToAlias[q.tableName]) > 0 {
		table += " AS " + q.QuoteIdentifier(q.mapTableToAlias[q.tableName])
	}
	return fmt.Sprintf("%s %s FROM %s %s", q.queryType, strings.Join(q.fields, ", "), table, tail)
}

func (q *QueryBuilder) buildInsert(where string) string {
	sql := fmt.Sprintf("INSERT INTO %s %s", q.tableNameEscaped, q.insert)
	if len(q.returningFields) > 0 {
		for i, v := range q.returningFields {
			q.returningFields[i] = q.QuoteIdentifier(v)
		}
		tail := strings.Join(q.returningFields, ", ")
		sql += " RETURNING " + tail
	}
	return sql
}

func (q *QueryBuilder) buildUpdate(tail string) string {
	sql := fmt.Sprintf("UPDATE %s SET %s %s", q.tableNameEscaped, strings.Join(q.update, ", "), tail)
	if len(q.returningFields) > 0 {
		for i, v := range q.returningFields {
			q.returningFields[i] = q.QuoteIdentifier(v)
		}
		tail := strings.Join(q.returningFields, ", ")
		sql += " RETURNING " + tail
	}
	return sql
}

func sqlLimitString(offset int, limit int) string {
	if limit > 0 {
		if offset > 0 {
			return fmt.Sprintf("LIMIT %d,%d", offset, limit)
		}
		return fmt.Sprintf("LIMIT %d", limit)
	}
	return fmt.Sprintf("LIMIT %d,18446744073709551615", offset)
}

func sqlTailString(list ...string) string {
	var ret string
	for _, s := range list {
		if len(s) > 0 {
			ret += " " + strings.Trim(s, " ")
		}
	}
	return strings.Trim(ret, " ")
}
