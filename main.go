package main

import (
	"go-gin-postgre/database"
	"go-gin-postgre/routers"
)

func main() {
	database.DBMigrate(database.DB)
	routers.StartServer().Run(":8080")
}