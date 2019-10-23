package configloader

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type LoaderFunction = func(data []byte, v interface{}) error

var Extensions = make(map[string]LoaderFunction)

const DefaultExt = "json"

func init() {
	RegisterExtensionHandler("json", jsonLoader)
}

func jsonLoader(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func formatExt(ext string) string {
	if ext[0] == '.' {
		return ext[1:]
	}
	return ext
}

// 注册文件扩展名对应的解析器
func RegisterExtensionHandler(ext string, loader LoaderFunction) {
	Extensions[formatExt(ext)] = loader
}

// 指定扩展名和数据，加载配置
func Load(ext string, data []byte, v interface{}) error {
	loader, ok := Extensions[formatExt(ext)]
	if !ok {
		loader, ok = Extensions[DefaultExt]
		if !ok {
			panic("unexpected error, missing default loader")
		}
	}
	return loader(data, v)
}

// 加载指定文件的配置
func LoadFile(file string, v interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	ext := filepath.Ext(file)
	return LoadReader(ext, f, v)
}

// 加载指定文件的配置，如果出错则panic
func MustLoadFile(file string, v interface{}) {
	if err := LoadFile(file, v); err != nil {
		panic(err)
	}
}

// 指定扩展名和Reader，加载配置
func LoadReader(ext string, reader io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return Load(ext, data, v)
}
