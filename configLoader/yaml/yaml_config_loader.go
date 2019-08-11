package yaml

import (
	"github.com/go-yaml/yaml"

	"github.com/leizongmin/go-common-libs/configLoader"
)

// 支持yaml格式的配置文件
func init() {
	configLoader.RegisterExtensionHandler("yaml", yamlLoader)
	configLoader.RegisterExtensionHandler("yml", yamlLoader)
}

func yamlLoader(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
