package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"my-tracking-list-backend/infrastructure/config"
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

	config.InitIoCManager(r)

	log.Fatal(r.Run(fmt.Sprintf(":%s", os.Getenv("MY_TRACKING_LIST_PORT"))))
}
