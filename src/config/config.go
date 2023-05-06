package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"port"`
	DBConnectionString string `mapstructure:"connection_string"`
	Environment        string `mapstructure:"environment"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/dist/config") // <- config path for docker file
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Configurations Loaded Successfully" + fmt.Sprintf("%+v", AppConfig))
}
