package randUtils

import (
	"math/rand"
	"time"
)

var source rand.Source
var seed int64
var r *rand.Rand

func init() {
	SetSeed(time.Now().UnixNano())
}

// 设置随机种子
func SetSeed(s int64) {
	seed = s
	source = rand.NewSource(seed)
	r = rand.New(source)
	rand.Seed(seed)
}

func Get() *rand.Rand {
	return r
}

// 获取float64随机数
func Float64() float64 {
	return r.Float64()
}

// 获取float32随机数
func Float32() float32 {
	return r.Float32()
}

// 获取int随机数
func Intn(n int) int {
	return r.Intn(n)
}

// 获取int63随机数
func Int63n(n int64) int64 {
	return r.Int63n(n)
}
