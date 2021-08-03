package database

import (
	"fmt"
	"main/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
)

func InitDatabase() *gorm.DB {
	fmt.Println("Initializing database...")
	dns := ""
	if os.Getenv("env") == "production" {
		dns = "host=db user=quickventory password=123456 dbname=quickventory port=5432 sslmode=disable TimeZone=America/Toronto"
	} else {
		dns = "host=0.0.0.0 user=quickventory password=123456 dbname=quickventory port=5432 sslmode=disable TimeZone=America/Toronto"
	}
	var err error
	dbInstance, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Connection to database failed")
	}

	err = migrate(dbInstance)
	if err != nil {
		fmt.Println("Error migrating database: ", err)
	} else {
		fmt.Println("Database migrated successfully")
	}

	return dbInstance
}

func migrate(dbInstance *gorm.DB) error {
	if os.Getenv("env") != "production" {
		fmt.Println("Migrating database...")
		return dbInstance.AutoMigrate(models.User{})
	}
	return nil
}
