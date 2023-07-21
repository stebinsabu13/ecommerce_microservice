package config

import "github.com/spf13/viper"

type Config struct {
	Db_port        string `mapstructure:"DB_PORT"`
	Db_host        string `mapstructure:"DB_HOST"`
	Db_username    string `mapstructure:"DB_USER"`
	Db_password    string `mapstructure:"DB_PASSWORD"`
	Db_name        string `mapstructure:"DB_NAME"`
	AuthService    string `mapstructure:"AUTH_SRV"`
	ProductService string `mapstructure:"PRODUCT_SRV"`
	// UserService     string `mapstructure:"USER_SRV"`
	CartServiceUrl string `mapstructure:"CART_SRV"`
	// OrderServiceUrl string `mapstructure:"ORDER_SRV"`
	Port           string `mapstructure:"PORT"`
	RAZORPAYKEY    string `mapstructure:"RAZORPAY_KEY"`
	RAZORPAYSECRET string `mapstructure:"RAZORPAY_SECRET"`
}

var config *Config

var envs = []string{"DB_PORT", "DB_HOST",
	"DB_USER", "DB_PASSWORD", "DB_NAME",
	"PRODUCT_SRV", "CART_SRV", "PORT", "AUTH_SRV",
	"RAZORPAY_KEY", "RAZORPAY_SECRET",
}

func LoadConfig() (*Config, error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return nil, err
		}
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func GetCofig() *Config {
	return config
}
