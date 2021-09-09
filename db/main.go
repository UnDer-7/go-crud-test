package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Erro ao carregar variaveis de ambiente: %v", err)
	}

	databaseURL := os.Getenv("GO_CRUD_TEST_DATABASE_FULL_URL")
	databaseDriver := os.Getenv("GO_CRUD_TEST_DATABASE_DRIVER")

	db, err := sql.Open(
		databaseDriver,
		databaseURL,
	)

	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Nao foi possivel pingar o banco de dados: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Nao foi possivel criar instancia do banco de dados %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres",
		driver,
	)
	defer m.Close()

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao sincrionizar o banco de dados: %v", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		println("Nenhuma migration rodada, Banco de Dados ja esta sincronizado")
		os.Exit(0)
	}

	fmt.Println("Migrations rodadas, Banco de Dados Migrado")
	os.Exit(0)
}
