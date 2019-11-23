package pgsqlutil

import (
	"database/sql/driver"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/lib/pq"

	"github.com/leizongmin/go-common-libs/sqlutil"
)

type ConnectionOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string // disable, require, verify-ca, verify-full
	Params   map[string]string
}

func (opts ConnectionOptions) BuildDataSourceString() string {
	if opts.Host == "" {
		opts.Host = "127.0.0.1"
	}
	if opts.Port == 0 {
		opts.Port = 3306
	}
	if opts.User == "" {
		opts.User = "root"
	}
	str := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?",
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
	)
	appendQueryString := func(s string) {
		if str[len(str)-1:] == "?" {
			str += s
		} else {
			str += "&" + s
		}
	}
	if opts.SSLMode != "" {
		appendQueryString("sslmode=" + opts.SSLMode)
	}
	for k, v := range opts.Params {
		appendQueryString(k + "=" + url.QueryEscape(v))
	}
	return str
}

func Table(tableName string) *sqlutil.QueryBuilder {
	return sqlutil.NewEmptyQuery().Init(pq.QuoteIdentifier, quoteLiteral).Table(tableName)
}

func Custom(query string, args ...driver.Value) string {
	if len(args) < 1 {
		return strings.Trim(query, " ")
	}
	ret, err := sqlutil.InterpolateParams(query, args, sqlutil.GetDefaultLocation(), quoteLiteral)
	if err != nil {
		log.Println(err)
		return query
	}
	return strings.Trim(ret, " ")
}

func quoteLiteral(buf []byte, v string) []byte {
	v2 := []byte(pq.QuoteLiteral(v))
	buf = append(buf, v2...)
	return buf
}
