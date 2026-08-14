package libs

import (
	"github.com/spf13/viper"
)

type Paths struct {
	Path_mega string `mapstructure:"path_mega"`
}

type ConfigApp struct {
	Paths Paths `mapstructure:"paths"`
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

	return &config, nil
}
