package cliargs

import (
	"regexp"

	"github.com/leizongmin/go/textutil"
)

type CliArgs struct {
	RawArgs []string              // 原始参数
	Options map[string]OptionItem // 选项参数
	Args    []string              // 参数
}

type OptionItem struct {
	Key   string // 键
	Value string // 值
	Raw   string // 原始值
}

// 解析命令行参数
// 支持解析 -name=value, --name=value, -name, --name 这种形式
func Parse(rawArgs []string) *CliArgs {
	return (&CliArgs{RawArgs: rawArgs, Options: map[string]OptionItem{}}).parse()
}

func (a *CliArgs) parse() *CliArgs {
	reg1, err := regexp.Compile("^\\-\\-?([\\w\\-_]+)=(.*)$")
	if err != nil {
		panic(err)
	}
	reg2, err := regexp.Compile("^\\-\\-?([\\w\\-_]+)$")
	if err != nil {
		panic(err)
	}

	for _, s := range a.RawArgs {
		ret1 := reg1.FindStringSubmatch(s)
		if len(ret1) > 0 {
			a.Options[ret1[1]] = OptionItem{
				Key:   ret1[1],
				Value: ret1[2],
				Raw:   s,
			}
			continue
		}
		ret2 := reg2.FindStringSubmatch(s)
		if len(ret2) > 0 {
			a.Options[ret2[1]] = OptionItem{
				Key:   ret2[1],
				Value: "",
				Raw:   s,
			}
			continue
		}
		a.Args = append(a.Args, s)
	}
	return a
}

// 是否包含指定选项
func (a *CliArgs) HasOption(name string) bool {
	_, ok := a.Options[name]
	return ok
}

// 获取指定选项
func (a *CliArgs) GetOption(name string) OptionItem {
	return a.Options[name]
}

// 获取指定选项，如果不存在则返回默认值
func (a *CliArgs) GetOptionOrDefault(name string, defaultValue string) OptionItem {
	if _, ok := a.Options[name]; !ok {
		return OptionItem{Key: name, Value: defaultValue}
	}
	return a.Options[name]
}

// 选项的数量
func (a *CliArgs) OptionsCount() int {
	return len(a.Options)
}

// 遍历所有选项
func (a *CliArgs) ForEachOptions(handler func(item OptionItem)) {
	for _, item := range a.Options {
		handler(item)
	}
}

// 参数数量
func (a *CliArgs) ArgsCount() int {
	return len(a.Args)
}

// 取指定索引的参数，如果不存在则返回空字符串
func (a *CliArgs) GetArg(i int) string {
	if i >= len(a.Args) {
		return ""
	}
	return a.Args[i]
}

// 遍历所有参数
func (a *CliArgs) ForEachArgs(handler func(item string)) {
	for _, item := range a.Args {
		handler(item)
	}
}

// 取指定索引位置开始的参数，如果不存在则返回空的Args
func (a *CliArgs) SubArgs(i int) *CliArgs {
	if i >= len(a.Args) {
		return &CliArgs{
			RawArgs: a.RawArgs,
			Options: a.Options,
			Args:    []string{},
		}
	}
	return &CliArgs{
		RawArgs: a.RawArgs,
		Options: a.Options,
		Args:    a.Args[i:],
	}
}

func (o OptionItem) TryParseInt(defaultValue int) int {
	return textutil.TryParseInt(o.Value, defaultValue)
}

func (o OptionItem) TryParseInt64(defaultValue int64) int64 {
	return textutil.TryParseInt64(o.Value, defaultValue)
}

func (o OptionItem) TryParseUint(defaultValue uint) uint {
	return textutil.TryParseUint(o.Value, defaultValue)
}

func (o OptionItem) TryParseUint64(defaultValue uint64) uint64 {
	return textutil.TryParseUint64(o.Value, defaultValue)
}

func (o OptionItem) TryParseFloat32(defaultValue float32) float32 {
	return textutil.TryParseFloat32(o.Value, defaultValue)
}

func (o OptionItem) TryParseFloat64(defaultValue float64) float64 {
	return textutil.TryParseFloat64(o.Value, defaultValue)
}

func (o OptionItem) TryParseBool(defaultValue bool) bool {
	return textutil.TryParseBool(o.Value, defaultValue)
}
