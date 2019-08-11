package jwtUtils

import (
	"fmt"
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	secret string
}

type Value = map[string]interface{}

func New(secret string) *JWT {
	return &JWT{secret}
}

// 签名
func (j *JWT) Sign(data Value) (string, error) {
	calims := jwt2.MapClaims{
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
	}
	for k, v := range data {
		calims[k] = v
	}
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, calims)
	return token.SignedString([]byte(j.secret))
}

// 签名，如果出错则panic
func (j *JWT) MustSign(data Value) string {
	str, err := j.Sign(data)
	if err != nil {
		panic(err)
	}
	return str
}

// 解析签名
func (j *JWT) UnSign(signedStr string) (Value, error) {
	token, err := jwt2.Parse(signedStr, func(token *jwt2.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("faield#1")
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("failed#2")
	}
	data, ok := token.Claims.(jwt2.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed#3")
	}
	return data, nil
}

// 解析签名，如果出错则panic
func (j *JWT) MustUnSign(signedStr string) Value {
	v, err := j.UnSign(signedStr)
	if err != nil {
		panic(err)
	}
	return v
}
