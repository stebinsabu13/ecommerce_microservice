package config

import "github.com/spf13/viper"

type Config struct {
	Db_port        string `mapstructure:"DB_PORT"`
	Db_host        string `mapstructure:"DB_HOST"`
	Db_username    string `mapstructure:"DB_USER"`
	Db_password    string `mapstructure:"DB_PASSWORD"`
	Db_name        string `mapstructure:"DB_NAME"`
	Port           string `mapstructure:"PORT"`
	Cart_service   string `mapstructure:"CART_SRV"`
	Jwt_secret_key string `mapstructure:"JWT_SECRET_KEY"`
	Auth_token     string `mapstructure:"AUTH_TOKEN"`
	Account_sid    string `mapstructure:"ACCOUNT_SID"`
	Service_sid    string `mapstructure:"SERVICE_SID"`
}

var envs = []string{"DB_PORT", "DB_HOST",
	"DB_USER", "DB_PASSWORD", "DB_NAME", "PORT", "JWT_SECRET_KEY",
	"AUTH_TOKEN", "ACCOUNT_SID", "SERVICE_SID", "CART_SRV",
}

var config *Config

func LoadConfig() (*Config, error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return nil, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetCofig() *Config {
	return config
}
