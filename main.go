package main

import (
	"gin_jwt/config"
	"gin_jwt/routes"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
    // for load godotenv
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // programmatically set swagger info
    

    // database connection
    db := config.ConnectDatabase()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()

    // router
    r := routes.SetupRouter(db)
    r.Run("localhost:8080")
}