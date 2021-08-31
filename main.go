package main

import (
	"crud-test/infrastructure/ioc"
	"crud-test/infrastructure/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(middleware.HandlePanicRecovery))
	r.Use(middleware.HandleAppError)

	ioc.InitContainerManager(r)

	log.Fatal(r.Run(":8080"))
}
