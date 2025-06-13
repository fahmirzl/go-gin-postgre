package main

import (
	"fmt"
	"go-gin-postgre/routers"
	"os"
)

func main() {
	// database.DBMigrate(database.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Running on port:", port)
	routers.StartServer().Run(":" + port)
}