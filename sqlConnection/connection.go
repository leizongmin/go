package sqlConnection

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"sync/atomic"
)

var queryCounter int64

func incrQueueCounter() {
	atomic.AddInt64(&queryCounter, 1)
}

type Options struct {
	Host       string
	Port       int
	User       string
	Password   string
	Database   string
	Charset    string
	Timezone   string // +8:00
	ParseTime  bool
	AutoCommit bool
	Params     map[string]string
}

// 构建连接字符串
func BuildDataSourceString(opts Options) string {
	if opts.Host == "" {
		opts.Host = "127.0.0.1"
	}
	if opts.Port == 0 {
		opts.Port = 3306
	}
	if opts.User == "" {
		opts.User = "root"
	}
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
	)
	if opts.ParseTime {
		str += "&parseTime=true&loc=Local"
	}
	if opts.Charset != "" {
		str += "&charset=" + opts.Charset
	}
	if opts.Timezone != "" {
		str += "&time_zone=" + url.QueryEscape("'"+opts.Timezone+"'")
	}
	if opts.AutoCommit {
		str += "&autocommit=true"
	}
	for k, v := range opts.Params {
		str += "&" + k + "=" + url.QueryEscape(v)
	}
	return str
}

type DBBase interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type DB interface {
	DBBase
	MustBegin() Tx
}

type Tx interface {
	DBBase
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
func FindOne(tx DBBase, dest interface{}, query string, args ...interface{}) (success bool) {
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
func FindMany(tx DBBase, dest interface{}, query string, args ...interface{}) (success bool) {
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
func InsertOne(tx DBBase, query string, args ...interface{}) (insertId int) {
	incrQueueCounter()
	var err error
	var res sql.Result
	debugf("#%d InsertOne: %s %+v", queryCounter, query, args)
	res, err = tx.Exec(query, args...)
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
		return 0
	}
	id, err := res.LastInsertId()
	if err != nil {
		warningf("#%d InsertOne failed: %s => %s %+v", queryCounter, err, query, args)
	}
	insertId = int(id)
	debugf("#%d InsertOne: insertId=%d", queryCounter, insertId)
	return insertId
}

// 更新多条数据
func UpdateMany(tx DBBase, query string, args ...interface{}) (rowsAffected int) {
	incrQueueCounter()
	debugf("#%d UpdateMany: %s %+v", queryCounter, query, args)
	res, err := tx.Exec(query, args...)
	if err != nil {
		warningf("UpdateMany failed: %s => %s %+v", err, query, args)
		return 0
	}
	rows, err := res.RowsAffected()
	if err != nil {
		warningf("UpdateMany failed: %s => %s %+v", err, query, args)
	}
	rowsAffected = int(rows)
	debugf("#%d UpdateMany: rowsAffected=%d", queryCounter, rowsAffected)
	return rowsAffected
}

// 更新一条数据
func UpdateOne(tx DBBase, query string, args ...interface{}) (rowsAffected int) {
	incrQueueCounter()
	rowsAffected = UpdateMany(tx, query+" LIMIT 1", args...)
	return rowsAffected
}

type queryCountRow struct {
	Count int `db:"count"`
}

// 查询记录数量，需要 SELECT count(*) AS count FROM ... 这样的格式
func FindCount(tx DBBase, query string, args ...interface{}) (count int, success bool) {
	row := new(queryCountRow)
	ok := FindOne(tx, row, query, args...)
	if ok {
		return row.Count, true
	}
	return 0, false
}
