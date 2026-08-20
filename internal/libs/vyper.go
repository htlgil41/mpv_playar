package libs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Paths struct {
	Path_mega      string `mapstructure:"path_mega"`
	Path_servermpv string `mapstructure:"path_server_mpv"`
}

type Server struct {
	Port   int64  `mapstructure:"port"`
	DBLite string `mapstructure:"dblite"`
}

type MailsSendersNotifier struct {
	To          string `mapstructure:"to_gmail"`
	Password_to string `mapstructure:"password_togmail"`
	From        string `mapstructure:"from_gmail"`
}

type App struct {
	Sucursal string `mapstructure:"sucursal"`
}

type ConfigApp struct {
	App                  App                   `mapstructure:"app"`
	Paths                Paths                 `mapstructure:"paths"`
	Server               Server                `mapstructure:"server"`
	MailsSendersNotifier *MailsSendersNotifier `mapstructure:"mailssendersnotifier"`
}

func LoadConfigWithVyper(path string) (*ConfigApp, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err_read := viper.ReadInConfig(); err_read != nil {
		return nil, err_read
	}

	var config ConfigApp
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	fmt.Println(config)
	return &config, nil
}
