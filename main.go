package main

import (
	"crud-test/infrastructure/ioc"
	"crud-test/infrastructure/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar variaveis de ambiente: %v", err)
	}

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(middleware.HandlePanicRecovery))
	r.Use(middleware.HandleAppError)

	ioc.InitContainerManager(r)

	log.Fatal(r.Run(":8080"))
}
