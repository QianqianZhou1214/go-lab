package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnvironment() {
	if err := godotenv.Load("app.env"); err != nil {
		logrus.Fatal("Error loading .env file")
	}
}
