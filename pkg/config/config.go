package config

import (
	"log"
	"os"
)

type Configure struct {
	PORT         string `json:"PORT"`
	ADMINPORT    string `json:"ADMINPORT"`
	BSERVICEPORT string `json:"BSERVICEPORT"`
	SECRETKEY    string `json:"SECRETKEY"`
	REDISHOST    string `json:"REDISHOST"`
}

func LoadConfigure() (*Configure, error) {
	var cfg Configure

	// Read from environment variables
	//if err := godotenv.Load("../../.env"); err != nil {
	//	os.Exit(1)
	//}

	cfg.PORT = os.Getenv("PORT")
	cfg.ADMINPORT = os.Getenv("ADMINPORT")
	cfg.BSERVICEPORT = os.Getenv("BSERVICEPORT")
	cfg.SECRETKEY = os.Getenv("SECRETKEY")
	cfg.REDISHOST = os.Getenv("REDISHOST")

	log.Println("api-gateway-service env -> ", cfg)

	return &cfg, nil
}
