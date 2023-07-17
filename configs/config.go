package configs

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Server Server
	Store  Store
}

type Server struct {
	Address string `json:"address"`
}

type Store struct {
	DB  PSQL
	RDB Redis
}

type PSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Name     string `json:"dbname"`
	SSL      string `json:"sslmode"`
}

type Redis struct {
	Address string `json:"address"`
}

func New() (Configs, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	cfg := Configs{
		Server: Server{
			Address: viper.GetString("server.address"),
		},
		Store: Store{
			DB: PSQL{
				Host:     viper.GetString("psql.localhost"),
				Port:     viper.GetString("psql.port"),
				Username: viper.GetString("psql.username"),
				Name:     viper.GetString("psql.dbname"),
				SSL:      viper.GetString("psql.sslmode"),
			},
			RDB: Redis{
				Address: viper.GetString("redis.address"),
			},
		},
	}

	return cfg, nil
}
