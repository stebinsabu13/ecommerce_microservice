package config

import "github.com/spf13/viper"

type Config struct {
	AuthService     string `mapstructure:"AUTH_SRV"`
	ProductService  string `mapstructure:"PRODUCT_SRV"`
	CartServiceUrl  string `mapstructure:"CART_SRV"`
	OrderServiceUrl string `mapstructure:"ORDER_SRV"`
	Port            string `mapstructure:"PORT"`
}

var envs = []string{"AUTH_SRV",
	"PRODUCT_SRV", "CART_SRV", "ORDER_SRV", "PORT",
}

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			return
		}
	}
	err = viper.Unmarshal(&config)

	return
}
