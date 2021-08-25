package main

import (
	"alta-store-project/config"
	"alta-store-project/routes"
	"fmt"
)

func main() {
	config.InitDB()
	config.InitMigrate()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
	fmt.Println("test midd")
}
