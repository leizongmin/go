package cliargs

import (
	"regexp"
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
func Parse(rawArgs []string) *CliArgs {
	return (&CliArgs{RawArgs: rawArgs, Options: map[string]OptionItem{}}).parse()
}

func (a *CliArgs) parse() *CliArgs {
	reg, err := regexp.Compile("^\\-\\-?(\\w+)=(.*)$")
	if err != nil {
		panic(err)
	}
	for _, s := range a.RawArgs {
		ret := reg.FindStringSubmatch(s)
		if len(ret) > 0 {
			a.Options[ret[1]] = OptionItem{
				Key:   ret[1],
				Value: ret[2],
				Raw:   s,
			}
		} else {
			a.Args = append(a.Args, s)
		}
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

// 取指定索引的参数
func (a *CliArgs) GetArg(i int) string {
	return a.Args[i]
}

// 遍历所有参数
func (a *CliArgs) ForEachArgs(handler func(item string)) {
	for _, item := range a.Args {
		handler(item)
	}
}

// 取指定索引位置开始的参数
func (a *CliArgs) SubArgs(i int) *CliArgs {
	return &CliArgs{
		RawArgs: a.RawArgs,
		Options: a.Options,
		Args:    a.Args[i:],
	}
}
