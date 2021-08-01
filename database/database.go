package database

import (
	"fmt"
	"main/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *DB

type DB struct {
	*gorm.DB // or what database you want like *mongo.Client
}

func GetDB() *DB {
	if db == nil {
		db.DB = InitDatabase()
	}
	return db
}

func InitDatabase() *gorm.DB {
	fmt.Println("Initializing database...")
	dns := ""
	if os.Getenv("env") == "production" {
		dns = "host=db user=quickventory password=123456 dbname=quickventory port=5432 sslmode=disable TimeZone=America/Toronto"
	} else {
		dns = "host=0.0.0.0 user=quickventory password=123456 dbname=quickventory port=5432 sslmode=disable TimeZone=America/Toronto"
	}
	var err error
	db.DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Connection to database failed")
	}

	err = migrate(db.DB)
	if err != nil {
		fmt.Println("Error migrating database: ", err)
	} else {
		fmt.Println("Database migrated successfully")
	}

	return db.DB
}

func migrate(dbInstance *gorm.DB) error {
	if os.Getenv("env") != "production" {
		fmt.Println("Migrating database...")
		return dbInstance.AutoMigrate(models.User{})
	}
	return nil
}
