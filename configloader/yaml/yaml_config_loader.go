package yaml

import (
	"gopkg.in/yaml.v2"

	"github.com/leizongmin/go/configloader"
)

// 支持yaml格式的配置文件
func init() {
	configloader.RegisterExtensionHandler("yaml", yamlLoader)
	configloader.RegisterExtensionHandler("yml", yamlLoader)
}

func yamlLoader(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
