package main

import (
	"go-gin-postgre/database"
	"go-gin-postgre/routers"
	"os"
)

func main() {
	database.DBMigrate(database.DB)
	routers.StartServer().Run(":" + os.Getenv("PORT"))
}