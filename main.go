package main

import "go-gin-postgre/routers"

func main() {
	routers.StartServer().Run(":8080")
}