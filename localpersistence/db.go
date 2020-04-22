package localpersistence

import (
	"log"
	"os"
	"path"

	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type DBOptions = opt.Options

type DB = leveldb.DB

type Options struct {
	DB      *DBOptions
	Encoder EncoderFunc
	Decoder DecoderFunc
}

type EncoderFunc func(value interface{}) ([]byte, error)
type DecoderFunc func([]byte, interface{}) error

var DefaultEncoder = func(value interface{}) ([]byte, error) {
	return jsoniter.Marshal(value)
}

var DefaultDecoder = func(b []byte, value interface{}) error {
	return jsoniter.Unmarshal(b, value)
}

// 填充默认配置
func fillDefaultOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{}
	}
	if opts.DB == nil {
		opts.DB = &DBOptions{}
	}
	if opts.Encoder == nil {
		opts.Encoder = DefaultEncoder
	}
	if opts.Decoder == nil {
		opts.Decoder = DefaultDecoder
	}
	return opts
}

// 打开leveldb数据库
func openDb(file string, opts *DBOptions) (*DB, error) {
	if err := os.MkdirAll(path.Dir(file), os.ModePerm); err != nil {
		log.Println(err)
	}
	return leveldb.OpenFile(file, opts)
}
