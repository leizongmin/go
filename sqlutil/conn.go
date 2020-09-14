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
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
}

type AbstractDB interface {
	AbstractDBBase
	MustBegin() *Tx
	Beginx() (*Tx, error)
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

// 执行查询，无返回结果
func Exec(tx AbstractDBBase, query string, args ...interface{}) (rowsAffected int64, success bool) {
	incrQueueCounter()
	debugf("#%d Exec: %s %+v", queryCounter, query, args)
	res, err := tx.Exec(query, args...)
	if err != nil {
		warningf("Exec failed: %s => %s %+v", err, query, args)
		return 0, false
	}
	rows, err := res.RowsAffected()
	if err != nil {
		warningf("Exec failed: %s => %s %+v", err, query, args)
	}
	rowsAffected = rows
	debugf("#%d Exec: rowsAffected=%d", queryCounter, rowsAffected)
	return rowsAffected, true
}

// 执行插入记录查询，返回最后插入ID
func ExecInsert(tx AbstractDBBase, query string, args ...interface{}) (lastInsertId int64, success bool) {
	incrQueueCounter()
	var err error
	var res sql.Result
	debugf("#%d ExecInsert: %s %+v", queryCounter, query, args)
	res, err = tx.Exec(query, args...)
	if err != nil {
		warningf("#%d ExecInsert failed: %s => %s %+v", queryCounter, err, query, args)
		return 0, false
	}
	id, err := res.LastInsertId()
	if err != nil {
		warningf("#%d ExecInsert failed: %s => %s %+v", queryCounter, err, query, args)
	}
	lastInsertId = id
	debugf("#%d ExecInsert: insertId=%d", queryCounter, lastInsertId)
	return lastInsertId, true
}

// 执行查询，有一行返回结果
func QueryOne(tx AbstractDBBase, dest interface{}, query string, args ...interface{}) (success bool) {
	incrQueueCounter()
	debugf("#%d QueryOne: %s %+v", queryCounter, query, args)
	err := tx.Get(dest, query, args...)
	if err != nil {
		if err != sql.ErrNoRows {
			warningf("#%d QueryOne failed: %s => %s %+v", queryCounter, err, query, args)
		}
		debugf("#%d QueryOne: success=false", queryCounter)
		return false
	}
	debugf("#%d QueryOne: success=true", queryCounter)
	return true
}

// 执行查询，有多行返回结果
func QueryMany(tx AbstractDBBase, dest interface{}, query string, args ...interface{}) (success bool) {
	incrQueueCounter()
	debugf("#%d QueryMany: %s %+v", queryCounter, query, args)
	err := tx.Select(dest, query, args...)
	if err != nil {
		if err != sql.ErrNoRows {
			warningf("#%d QueryMany failed: %s => %s %+v", queryCounter, err, query, args)
		}
		debugf("#%d QueryMany: success=false", queryCounter)
		return false
	}
	debugf("#%d QueryMany: success=true", queryCounter)
	return true
}

// 执行查询，有一行Map返回结果，所有字段值为[]byte或nil类型，可以转换为字符串
func QueryOneToMap(tx AbstractDBBase, query string, args ...interface{}) (row Row, success bool) {
	incrQueueCounter()
	debugf("#%d QueryOneToMap: %s %+v", queryCounter, query, args)
	result, err := tx.Queryx(query, args...)
	defer result.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			warningf("#%d QueryOneToMap failed: %s => %s %+v", queryCounter, err, query, args)
		}
		debugf("#%d QueryOneToMap: success=false", queryCounter)
		return nil, false
	}
	if !result.Next() {
		debugf("#%d QueryOneToMap: success=false", queryCounter)
		return nil, false
	}
	debugf("#%d QueryOneToMap: success=true", queryCounter)
	row = make(Row)
	if err := result.MapScan(row); err != nil {
		debugf("#%d QueryOneToMap: success=false, %s", queryCounter, err)
		return nil, false
	}
	return row, true
}

// 执行查询，有多行Map返回结果，所有字段值为[]byte或nil类型，可以转换为字符串
func QueryManyToMap(tx AbstractDBBase, query string, args ...interface{}) (rows []Row, success bool) {
	incrQueueCounter()
	debugf("#%d QueryManyToMap: %s %+v", queryCounter, query, args)
	result, err := tx.Queryx(query, args...)
	defer result.Close()
	if err != nil {
		if err != sql.ErrNoRows {
			warningf("#%d QueryManyToMap failed: %s => %s %+v", queryCounter, err, query, args)
		}
		debugf("#%d QueryManyToMap: success=false", queryCounter)
		return nil, false
	}
	//if !result.Next() {
	//	debugf("#%d QueryManyToMap: success=false", queryCounter)
	//	return nil, false
	//}
	debugf("#%d QueryManyToMap: success=true", queryCounter)
	rows = make([]Row, 0)
	for result.Next() {
		row := make(Row)
		if err := result.MapScan(row); err != nil {
			debugf("#%d QueryManyToMap: success=false, %s", queryCounter, err)
			return nil, false
		}
		rows = append(rows, row)
	}
	return rows, true
}

type QueryCountRow struct {
	Count int64 `db:"count"`
}

// 查询记录数量，需要 SELECT count(*) AS count FROM ... 这样的格式
func QueryCount(tx AbstractDBBase, query string, args ...interface{}) (count int64, success bool) {
	row := new(QueryCountRow)
	ok := QueryOne(tx, row, query, args...)
	if ok {
		return row.Count, true
	}
	return 0, false
}

// 事务执行多行SQL
func BatchExec(db AbstractDB, sqlList ...string) error {
	incrQueueCounter()
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	go func() {
		if err != nil && tx != nil {
			debugf("#%d BatchExec failed %s", queryCounter, err)
			if err2 := tx.Rollback(); err2 != nil {
				debugf("#%d BatchExec rollback failed %s", queryCounter, err2)
			}
		}
	}()
	for i, s := range sqlList {
		debugf("#%d BatchExec: [%d] %s", queryCounter, i, s)
		r, err := tx.Exec(s)
		if err != nil {
			return err
		}
		rowsAffected, _ := r.RowsAffected()
		lastInsertId, _ := r.LastInsertId()
		debugf("#%d BatchExec: success=true rowsAffected=%d lastInsertId=%d", queryCounter, rowsAffected, lastInsertId)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
