package utils

import (
	"StudentManager/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		middleware.Ln.Fatal(err)
	}
	viper.WatchConfig()
}
