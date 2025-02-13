package env

import "github.com/spf13/viper"

type Config struct {
	GinMode      string `mapstructure:"GIN_MODE"`
	GinPort      string `mapstructure:"GIN_PORT"`
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbName       string `mapstructure:"DB_NAME"`
	DbUsername   string `mapstructure:"DB_USERNAME"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	DbTz         string `mapstructure:"DB_TZ"`
	DbLogLevel   string `mapstructure:"DB_LOG_LEVEL"`
	DbMigrate    bool   `mapstructure:"DB_MIGRATE"`
	SaltPassword int    `mapstructure:"SALT_PASSWORD"`
	SecretKey    string `mapstructure:"SECRET_KEY"`
	ExpiredTime  int    `mapstructure:"EXPIRED_TIME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	GlobalEnv = config
	return
}
