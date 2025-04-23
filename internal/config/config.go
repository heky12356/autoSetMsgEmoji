package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config struct {
	Hostadd string
	GroupID int64
}

func Init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Config.Hostadd = os.Getenv("HOSTADD")
	Config.GroupID, _ = strconv.ParseInt(os.Getenv("GROUPID"), 10, 64)
}
