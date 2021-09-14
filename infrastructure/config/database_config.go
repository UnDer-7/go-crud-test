package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConfig() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("GO_CRUD_TEST_DATABASE_HOST"),
		os.Getenv("GO_CRUD_TEST_DATABASE_USER"),
		os.Getenv("GO_CRUD_TEST_DATABASE_PASSWORD"),
		os.Getenv("GO_CRUD_TEST_DATABASE_NAME"),
		os.Getenv("GO_CRUD_TEST_DATABASE_PORT"),
		os.Getenv("GO_CRUD_TEST_DATABASE_SSL_MODE"),
		os.Getenv("GO_CRUD_TEST_DATABASE_TIME_ZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect repository. Err: %s", err.Error()))
	}

	return db
}
