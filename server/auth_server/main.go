package main

import (
	"auth-service/database"
	"auth-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	routes.AuthRoutes(r)

	r.Run(":8080")
}
