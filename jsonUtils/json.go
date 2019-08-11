package jsonUtils

import (
	jsoniter "github.com/json-iterator/go"
)

// 解析JSON，如果出错则panic
func MustUnmarshal(data []byte, v interface{}) {
	if err := jsoniter.Unmarshal(data, v); err != nil {
		panic(err)
	}
}

// 序列化JSON，如果出错则panic
func MustMarshal(v interface{}) []byte {
	data, err := jsoniter.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}
