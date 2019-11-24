package sqlutil

import (
	"database/sql"
	"log"
	"sync/atomic"

	"github.com/jmoiron/sqlx"
)

var queryCounter int64

func incrQueueCounter() {
	atomic.AddInt64(&queryCounter, 1)
}

type DB = sqlx.DB
type Tx = sqlx.Tx

type ConnectionOptions interface {
	BuildDataSourceString() string
}

// 创建数据库连接
func OpenWithOptions(driverName string, opts ConnectionOptions) (*sqlx.DB, error) {
	return Open(driverName, opts.BuildDataSourceString())
}

// 创建数据库连接
func Open(driverName string, dataSourceName string) (*sqlx.DB, error) {
	debugf("Open: %s %s", driverName, dataSourceName)
	return sqlx.Open(driverName, dataSourceName)
}

// 创建数据库连接，如果失败则panic
func MustOpenWithOptions(driverName string, opts ConnectionOptions) *sqlx.DB {
	return MustOpen(driverName, opts.BuildDataSourceString())
}

// 创建数据库连接，如果失败则panic
func MustOpen(driverName string, dataSourceName string) *sqlx.DB {
	debugf("MustOpen: %s %s", driverName, dataSourceName)
	return sqlx.MustOpen(driverName, dataSourceName)
}

type AbstractDBBase interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type AbstractDB interface {
	AbstractDBBase
	MustBegin() AbstractTx
}

type AbstractTx interface {
	AbstractDBBase
	Rollback() error
	Commit() error
}

var isDebug = false

func EnableDebug() {
	isDebug = true
}

func DisableDebug() {
	isDebug = false
}

func debugf(format string, args ...interface{}) {
	if isDebug {
		log.Printf("DEBUG\t"+format, args...)
	}
}

func warningf(format string, args ...interface{}) {
	if isDebug {
		log.Printf("WARN\t"+format, args...)
	}
}

type Row = map[string]interface{}

// 查询一条数据
func FindOne(tx AbstractDBBase, dest interface{}, query string, args ...interface{}) (success bool) {
	incrQueueCounter()
	debugf("#%d FindOne: %s %+v", queryCounter, query, args)
	err := tx.Get(dest, query, args...)
	if err != nil {
		if err != sql.ErrNoRows {
			warningf("#%d FindOne failed: %s => %s %+v", queryCounter, err, query, args)
		}
		debugf("#%d FindMany: success=false", queryCounter)
		return false
	}
	debugf("#%d FindMany: success=true", queryCounter)
	return true
}

// 查询多条数据
func FindMany(tx AbstractDBBase, dest interface{}, query string, args ...interface{}) (success bool) {
	incrQueueCounter()
	debugf("%#d FindMany: %s %+v", queryCounter, query, args)
	err := tx.Select(dest, query, args...)
	if err != nil {
		if err != sql.ErrNoRows {
			warningf("#%d FindMany failed: %s => %s %+v", queryCounter, err, query, args)
		}
		debugf("#%d FindMany: success=false", queryCounter)
		return false
	}
	debugf("#%d FindMany: success=true", queryCounter)
	return true
}

// 插入一条数据
func InsertOne(tx AbstractDBBase, query string, args ...interface{}) (insertId int64, success bool) {
	incrQueueCounter()
	var err error
	var res sql.Result
	debugf("#%d InsertOne: %s %+v", queryCounter, query, args)
	res, err = tx.Exec(query, args...)
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
		return 0, false
	}
	id, err := res.LastInsertId()
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
	}
	insertId = id
	debugf("#%d InsertOne: insertId=%d", queryCounter, insertId)
	return insertId, true
}

// 插入一条数据，不返回insertId
func InsertOne2(tx AbstractDBBase, query string, args ...interface{}) (success bool) {
	incrQueueCounter()
	var err error
	debugf("#%d InsertOne: %s %+v", queryCounter, query, args)
	_, err = tx.Exec(query, args...)
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
		return false
	}
	return true
}

// 插入多条记录
func InsertMany(tx AbstractDBBase, query string, args ...interface{}) (lastInsertId int64, success bool) {
	incrQueueCounter()
	var err error
	var res sql.Result
	debugf("#%d InsertOne: %s %+v", queryCounter, query, args)
	res, err = tx.Exec(query, args...)
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
		return 0, false
	}
	id, err := res.LastInsertId()
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
	}
	lastInsertId = id
	debugf("#%d InsertOne: insertId=%d", queryCounter, lastInsertId)
	return lastInsertId, true
}

// 插入多条记录，不返回insertId
func InsertMany2(tx AbstractDBBase, query string, args ...interface{}) (success bool) {
	incrQueueCounter()
	var err error
	debugf("#%d InsertOne: %s %+v", queryCounter, query, args)
	_, err = tx.Exec(query, args...)
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
		return false
	}
	return true
}

// 更新多条数据
func UpdateMany(tx AbstractDBBase, query string, args ...interface{}) (rowsAffected int64, success bool) {
	incrQueueCounter()
	debugf("#%d UpdateMany: %s %+v", queryCounter, query, args)
	res, err := tx.Exec(query, args...)
	if err != nil {
		warningf("UpdateMany failed: %s => %s %+v", err, query, args)
		return 0, false
	}
	rows, err := res.RowsAffected()
	if err != nil {
		warningf("UpdateMany failed: %s => %s %+v", err, query, args)
	}
	rowsAffected = rows
	debugf("#%d UpdateMany: rowsAffected=%d", queryCounter, rowsAffected)
	return rowsAffected, true
}

// 更新一条数据
func UpdateOne(tx AbstractDBBase, query string, args ...interface{}) (rowsAffected int64, success bool) {
	incrQueueCounter()
	return UpdateMany(tx, query+" LIMIT 1", args...)
}

// 删除多条数据
func DeleteMany(tx AbstractDBBase, query string, args ...interface{}) (rowsAffected int64, success bool) {
	incrQueueCounter()
	return UpdateMany(tx, query, args...)
}

// 删除一条数据
func DeleteOne(tx AbstractDBBase, query string, args ...interface{}) (rowsAffected int64, success bool) {
	incrQueueCounter()
	return UpdateMany(tx, query+" LIMIT 1", args...)
}

type QueryCountRow struct {
	Count int64 `db:"count"`
}

// 查询记录数量，需要 SELECT count(*) AS count FROM ... 这样的格式
func FindCount(tx AbstractDBBase, query string, args ...interface{}) (count int64, success bool) {
	row := new(QueryCountRow)
	ok := FindOne(tx, row, query, args...)
	if ok {
		return row.Count, true
	}
	return 0, false
}
