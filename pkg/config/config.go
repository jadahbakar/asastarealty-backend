package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		// AppConfig is
		AppName          string `mapstructure:"APP_NAME"`
		AppURLGroup      string `mapstructure:"APP_URL_GROUP"`
		AppURLVersion    string `mapstructure:"APP_URL_VERSION"`
		AppLogFolder     string `mapstructure:"APP_LOG_FOLDER"`
		AppPort          int    `mapstructure:"APP_PORT"`
		AppPrefork       bool   `mapstructure:"APP_PREFORK"`
		AppCaseSensitive bool   `mapstructure:"APP_CASE_SENSITIVE"`
		AppReadTimeOut   int    `mapstructure:"APP_READ_TIMEOUT"`
		AppWriteTimeOut  int    `mapstructure:"APP_WRITE_TIMEOUT"`

		// DatabaseConfig is
		DbUrl       string `mapstructure:"DATABASE_URL"`
		DbTimezone  string `mapstructure:"DATABASE_TIMEZOME"`
		DbParseTime string `mapstructure:"DATABASE_PARSETIME"`
	}
)

func New() (config *Config, err error) {
	viper.SetConfigFile(`app.env`)
	// viper.AddConfigPath("./../../../")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
