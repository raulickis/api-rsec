package config

import (
	"os"
	"strconv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	ENV                      = os.Getenv("ENVIRONMENT")
	APP_PORT                 = os.Getenv("APP_PORT")
	PORT_EXTERNAL            = os.Getenv("PORT_EXTERNAL")
	PORT_INTERNAL            = os.Getenv("PORT_INTERNAL")
	DOCKER_HOST_IP           = os.Getenv("DOCKER_HOST_IP")
	DB_HOST                  = os.Getenv("DB_HOST")
	DB_NAME                  = os.Getenv("DB_NAME")
	DB_USER                  = os.Getenv("DB_USER")
	DB_PASS                  = os.Getenv("DB_PASS")
	DB_PORT, _               = strconv.Atoi(os.Getenv("DB_PORT"))
	TOKEN                    = os.Getenv("TOKEN")

)
