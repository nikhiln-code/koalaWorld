package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct{
	App struct{
		Port string
	}

	Cors struct{
		AllowedOrigins []string `mapstructure:"allowed_origins"`
	}

	Blockchain struct{
		NodeURL string `mapstructure:"node_url"`
		PrivateKey string `mapstructure:"private_key"`
	}

	Pinata struct{
		JwtKEY string `mapstructure:"jwt_key"`
	}
}


func LoadConfig()(*AppConfig, error){
	env := os.Getenv("APP_ENV")
	if(env ==""){
		env ="dev"  //default to dev if not set in enviornment
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err:= viper.ReadInConfig(); err !=nil{
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	var conf AppConfig
	if err := viper.Unmarshal(&conf); err != nil{
		return nil,fmt.Errorf("Unable to decode config into struct: %w", err)
	}

	return &conf, nil

}