package main

import (
	"go-gin-postgre/database"
	"go-gin-postgre/routers"
	"os"
)

func main() {
	database.DBMigrate(database.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routers.StartServer().Run(":" + port)
}