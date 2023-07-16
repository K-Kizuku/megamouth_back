package main

import (
	"fmt"
	"log"

	"megamouth/api/driver"
	"megamouth/api/utils/config"
	_ "megamouth/docs"
)

// @title   Megamouth
// @version  1.0
// @license.name Kizuku
// @BasePath /api/v1
// @securityDefinitions.basic BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadEnv()
	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", config.ApiPort))
}
