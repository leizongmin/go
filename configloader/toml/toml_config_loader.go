package toml

import (
	"github.com/BurntSushi/toml"

	"github.com/leizongmin/go/configloader"
)

// 支持toml格式的配置文件
func init() {
	configloader.RegisterExtensionHandler("toml", tomlLoader)
}

func tomlLoader(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}
