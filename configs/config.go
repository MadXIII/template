package configs

import (
	"encoding/json"
	"io"
	"os"
)

type Configs struct {
	Server Server
	PSQL   PSQL
}

type Server struct {
	Address string `json:"address"`
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
	file, err := os.Open("./configs/config.json")
	if err != nil {
		return Configs{}, err
	}

	bb, err := io.ReadAll(file)
	if err != nil {
		return Configs{}, err
	}

	var cfg Configs

	if err := json.Unmarshal(bb, &cfg); err != nil {
		return Configs{}, err
	}

	return cfg, nil
}
