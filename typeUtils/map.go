package typeUtils

import (
	"encoding/json"
	"fmt"
)

func StructToMap(data interface{}) (result map[string]interface{}, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	result = make(map[string]interface{})
	if err = json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func MustStructToMap(data interface{}) map[string]interface{} {
	if result, err := StructToMap(data); err != nil {
		panic(fmt.Sprintf("failed to cast %+v to map[string]interface{}: %s", data, err))
	} else {
		return result
	}
}
