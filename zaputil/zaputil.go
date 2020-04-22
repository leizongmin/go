package zaputil

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Development      bool                   `json:"development" yaml:"development" toml:"development"`                // 是否开发环境
	Level            string                 `json:"level" yaml:"level" toml:"level"`                                  // 日志等级，DEBUG,INFO,WARN,ERROR,FATAL
	Encoding         string                 `json:"encoding" yaml:"encoding" toml:"encoding"`                         // encoding = json; console
	OutputPaths      []string               `json:"outputPaths"  yaml:"outputPaths" toml:"outputPaths"`               // 日志文件输出路径列表，stdout和stderr内置
	ErrorOutputPaths []string               `json:"errorOutputPaths" yaml:"errorOutputPaths" toml:"errorOutputPaths"` // 错误日志文件输出路径列表，stdout和stderr内置
	InitialFields    map[string]interface{} `json:"initialFields" yaml:"initialFields" toml:"initialFields"`          // 默认字段，每条日志都包含
}

func Create(opts *Options) (*zap.Logger, error) {
	if opts == nil {
		opts = &Options{
			Development: true,
			Level:       "DEBUG",
			Encoding:    "console",
		}
	}
	levelMap := map[string]zapcore.Level{
		"DEBUG": zap.DebugLevel,
		"INFO":  zap.InfoLevel,
		"WARN":  zap.WarnLevel,
		"ERROR": zap.ErrorLevel,
		"FATAL": zap.FatalLevel,
	}
	level, ok := levelMap[strings.ToUpper(opts.Level)]
	if !ok {
		level = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	outputPaths := opts.OutputPaths
	errorOutputPaths := opts.ErrorOutputPaths
	if len(outputPaths) < 1 {
		outputPaths = []string{"stdout"}
	}
	if len(errorOutputPaths) < 1 {
		errorOutputPaths = []string{"stderr"}
	}

	config := zap.Config{
		Level:             atomicLevel,
		Development:       opts.Development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          opts.Encoding,
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       outputPaths,
		ErrorOutputPaths:  errorOutputPaths,
		InitialFields:     opts.InitialFields,
	}
	return config.Build()
}
