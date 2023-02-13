// Package config 配置操作
package config

import (
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var runtimeViper *viper.Viper

// ConfigFunc 返回配置信息的函数类型
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 缓存所有配置信息
var ConfigFuncs map[string]ConfigFunc = make(map[string]ConfigFunc)

// Load 加载配置文件
// 在调用 Load 之前需要调用 Add 函数添加配置信息
// 在调用 Load 之后才可以获取配置
func Load(env string) {
	// 初始化实例
	runtimeViper = viper.New()

	// 添加配置文件所在路径 . 代表程序的工作目录
	runtimeViper.AddConfigPath(".")

	// 设置配置文件类型
	runtimeViper.SetConfigType("env")
	// 添加环境变量前缀，用以与系统环境变量区分
	runtimeViper.SetEnvPrefix("appenv")
	// 读取环境变量
	runtimeViper.AutomaticEnv()

	// 加载环境变量
	loadEnv(env)
	// 加载配置信息
	loadConfig()
}

// loadEnv 加载环境变量
func loadEnv(suffix string) error {
	envFileName := ".env"
	if len(suffix) > 0 {
		filePath := envFileName + "." + suffix
		if _, err := os.Stat(filePath); err == nil {
			envFileName = filePath
		}
	}
	runtimeViper.SetConfigName(envFileName)
	if err := runtimeViper.ReadInConfig(); err != nil {
		return err
	}
	runtimeViper.WatchConfig()
	return nil
}

// loadConfig 加载所有配置项
func loadConfig() {
	for name, fn := range ConfigFuncs {
		runtimeViper.Set(name, fn())
	}
}

// Add 新增配置项
func Add(name string, fn ConfigFunc) {
	ConfigFuncs[name] = fn
}

// Env 获取配置项
func Env(name string, defaultValue ...interface{}) interface{} {
	return getValueDefault(name, defaultValue...)
}

// Get 获取配置项
func Get(name string, defaultValue ...interface{}) string {
	return cast.ToString(getValueDefault(name, defaultValue...))
}

// GetInt 获取 int 类型配置项
func GetInt(name string, defaultValue ...interface{}) int {
	return cast.ToInt(getValueDefault(name, defaultValue...))
}

// GetUint 获取 uint 类型配置项
func GetUint(name string, defaultValue ...interface{}) uint {
	return cast.ToUint(getValueDefault(name, defaultValue...))
}

// GetInt64 获取 int64 类型配置项
func GetInt64(name string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(getValueDefault(name, defaultValue...))
}

// GetFloat64 获取 float64 类型配置项
func GetFloat64(name string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(getValueDefault(name, defaultValue...))
}

// GetBool 获取 bool 类型配置项
func GetBool(name string, defaultValue ...interface{}) bool {
	return cast.ToBool(getValueDefault(name, defaultValue...))
}

// GetMapStringString 获取 map[string]string 类型配置项
func GetStringMapString(name string, defaultValue ...interface{}) map[string]string {
	return cast.ToStringMapString(getValueDefault(name, defaultValue...))
}

// getValueDefault 获取值处理默认值情况
func getValueDefault(name string, defaultValue ...interface{}) interface{} {
	if !runtimeViper.IsSet(name) && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return runtimeViper.Get(name)
}
