package toml

import (
	"github.com/BurntSushi/toml"

	"github.com/leizongmin/go-common-libs/configLoader"
)

// 支持toml格式的配置文件
func init() {
	configLoader.RegisterExtensionHandler("toml", tomlLoader)
}

func tomlLoader(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}
