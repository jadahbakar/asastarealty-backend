package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App *Application
		Db  *Database
	}

	Application struct {
		Name          string
		URLGroup      string
		URLVersion    string
		LogFolder     string
		Port          int
		Prefork       bool
		CaseSensitive bool
		ReadTimeOut   int
		WriteTimeOut  int
	}

	Database struct {
		Url       string
		Timezone  string
		ParseTime string
	}
)

func New(configPath string) (config *Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	configuration := Config{
		App: &Application{
			Name:          viper.GetString("APP_NAME"),
			URLGroup:      viper.GetString("APP_URL_GROUP"),
			URLVersion:    viper.GetString("APP_URL_VERSION"),
			LogFolder:     viper.GetString("APP_LOG_FOLDER"),
			Port:          viper.GetInt("APP_PORT"),
			Prefork:       viper.GetBool("APP_PREFORK"),
			CaseSensitive: viper.GetBool("APP_CASE_SENSITIVE"),
			ReadTimeOut:   viper.GetInt("APP_READ_TIMEOUT"),
			WriteTimeOut:  viper.GetInt("APP_WRITE_TIMEOUT"),
		},
		Db: &Database{
			Url:       viper.GetString("DATABASE_URL"),
			Timezone:  viper.GetString("DATABASE_TIMEZOME"),
			ParseTime: viper.GetString("DATABASE_PARSETIME"),
		},
	}
	return &configuration, err
}
