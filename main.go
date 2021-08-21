package main

import (
	"alta-store-project/config"
	"alta-store-project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
