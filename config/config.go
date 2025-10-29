package config

import (
	"errors"
	"fmt"
	"os"
	"rest-api/internals/model"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvironment() {
	if err := godotenv.Load("app.env"); err != nil {
		logrus.Fatal("Error loading .env file")
	}
}

func SetUpDatabase() (*gorm.DB, error) {
	dsn := os.Getenv(DatabaseUrl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}
	if err := RunAutoMigration(db); err != nil {
		return nil, fmt.Errorf("error running migrations: %v", err)
	}
	return db, nil
}

func RunAutoMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		errorMessage := fmt.Sprintf("error auto migrate db: %v", err)
		logrus.Error(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}
