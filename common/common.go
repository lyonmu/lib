package common

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// PathExists 判断文件路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

// LoadConfig 加载并监听配置文件，实现文件内容热加载
func LoadConfig(m string, conf any) error {
	viper.SetConfigFile(m)
	viper.WatchConfig()
	var resp error
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(conf); err != nil {
			resp = err
		}
	})
	if err := viper.ReadInConfig(); err != nil {
		resp = err
	}
	if err := viper.Unmarshal(conf); err != nil {
		resp = err
	}
	return resp
}
