package main

import (
	"crud-test/infrastructure/config"
	"crud-test/infrastructure/ioc"
	"crud-test/infrastructure/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		// todo: Usar lib de log
		fmt.Printf("\nNenhuma vareavel sera carregada do .env!!! \tErr: %v\n\n", err)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(config.CORSConfig())

	r.Use(gin.CustomRecovery(middleware.HandlePanicRecovery))
	r.Use(middleware.HandleAppError)

	ioc.InitContainerManager(r)

	log.Fatal(r.Run(fmt.Sprintf(":%s", os.Getenv("GO_CRUD_TEST_PORT"))))
}
