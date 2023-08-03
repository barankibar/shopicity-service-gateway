package main

import (
	"log"

	"github.com/barankibar/shopicity-svc-gateway/pkg/auth"
	"github.com/barankibar/shopicity-svc-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}
	r := gin.Default()

	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
