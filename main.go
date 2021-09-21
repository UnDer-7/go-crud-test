package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"my-tracking-list-backend/infrastructure/config"
	"my-tracking-list-backend/infrastructure/ioc"
	"my-tracking-list-backend/infrastructure/middleware"
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

	log.Fatal(r.Run(fmt.Sprintf(":%s", os.Getenv("MY_TRACKING_LIST_PORT"))))
}
