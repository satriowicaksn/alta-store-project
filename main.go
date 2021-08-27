package main

import (
	"alta-store-project/config"
	"alta-store-project/routes"
	"fmt"
)

func main() {

	// tambahan
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }
	// end

	config.InitDB()
	config.InitMigrate()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
	fmt.Println("test midd")
}
