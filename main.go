package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	InitContainerManager(r)

	log.Fatal(r.Run(":8080"))
}
