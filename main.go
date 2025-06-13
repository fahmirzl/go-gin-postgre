package main

import (
	"go-gin-postgre/routers"
	"os"
)

func main() {
	// database.DBMigrate(database.DB)
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable not set!")
	}
	routers.StartServer().Run(":" + port)
}
