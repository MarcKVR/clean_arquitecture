package database

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcKVR/clean_arquitecture/cmd/config"
	"github.com/MarcKVR/clean_arquitecture/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?TimeZone=%s",
		cfg.DBDriver, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.TimeZone)

	db_connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db_connection.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db_connection, nil
}

// func GetConnection() (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?TimeZone=%s",
// 		os.Getenv("DRIVER_APP"),
// 		os.Getenv("DATABASE_USER"),
// 		os.Getenv("DATABASE_PASSWORD"),
// 		os.Getenv("HOST_APP"),
// 		os.Getenv("DATABASE_PORT"),
// 		os.Getenv("DATABASE_NAME"),
// 		os.Getenv("TIMEZONE"),
// 	)

// 	db_connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		return nil, err
// 	}

// 	// if os.Getenv("DATABASE_DEBUG") == "true" {
// 	// 	db_connection = db_connection.Debug()
// 	// }

// 	// if os.Getenv("DATABASE_AUTO_MIGRATE") == "true" {
// 	// 	if err := db_connection.AutoMigrate(&domain.Payment{}); err != nil {
// 	// 		return nil, err
// 	// 	}

// 	// 	if err := db_connection.AutoMigrate(&domain.User{}); err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// }

// 	return db_connection, nil
// }

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	}
}
