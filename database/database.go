package database

import (
	"fmt"
	"main/models"
	"main/utils/env_utils"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDatabase() *gorm.DB {
	fmt.Println("Initializing database...")
	if err := godotenv.Load(".env"); err != nil {
		panic("Couldn't find .env file")
	}

	dns := ""
	dbUser := env_utils.FetchEnvOrPanic("POSTGRES_USER")
	dbPass := env_utils.FetchEnvOrPanic("POSTGRES_PASSWORD")
	dbName := env_utils.FetchEnvOrPanic("POSTGRES_DB")
	dbPort := env_utils.FetchEnvOrPanic("POSTGRES_PORT")
	intermediateDNS := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Toronto", dbUser, dbPass, dbName, dbPort)
	if os.Getenv("env") == "production" {
		dns = fmt.Sprintf("host=db %s", intermediateDNS)
	} else {
		dns = fmt.Sprintf("host=0.0.0.0 %s", intermediateDNS)
	}
	var err error
	Database, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Connection to database failed")
	}

	err = migrate(Database)
	if err != nil {
		fmt.Println("Error migrating database: ", err)
	} else {
		fmt.Println("Database migrated successfully")
	}

	return Database
}

func migrate(dbInstance *gorm.DB) error {
	if os.Getenv("env") != "production" {
		fmt.Println("Migrating database...")
		return dbInstance.AutoMigrate(models.User{}, models.AccessToken{})
	}
	return nil
}
