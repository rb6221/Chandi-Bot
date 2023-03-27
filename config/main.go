package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Token   string
	AppID   string
	GuildID string
)

func init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	UpdateConfigs()

	viper.OnConfigChange(
		func(e fsnotify.Event) {
			if e.Op == fsnotify.Write {
				UpdateConfigs()
			}
		},
	)

	viper.WatchConfig()

}

func UpdateConfigs() {
	Token = viper.GetString("token")
	AppID = viper.GetString("app_id")
	GuildID = viper.GetString("guild_id")
}
