package httpUtils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	{
		res, err := Request().GET("https://cnodejs.org").Send()
		assert.NoError(t, err)
		defer res.Close()
		fmt.Println(res.Status(), res.Header(), string(res.MustBody()))
	}
	{
		res, err := Request().GET("https://cnodejs.org/api/v1/topics").SetQuery("limit", "1").AcceptJSON().Send()
		assert.NoError(t, err)
		defer res.Close()
		data := make(map[string]interface{})
		err = res.JSON(&data)
		assert.NoError(t, err)
		fmt.Println(res.Status(), res.Header(), data)
	}
}
