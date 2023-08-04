package main

import (
	"log"

	"github.com/barankibar/shopicity-svc-gateway/pkg/auth"
	"github.com/barankibar/shopicity-svc-gateway/pkg/config"
	"github.com/barankibar/shopicity-svc-gateway/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	// ROUTER INSTANCE
	router := gin.Default()

	// MIDDLEWARES
	router.Use(middlewares.ErrorHandler)

	auth.RegisterRoutes(router, &c)

	router.Run(c.Port)
}
