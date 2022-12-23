package main

import (
	"curd-api/database"
	"curd-api/router"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

	}
	database.CreatConnectionTodb()

	r := router.InitRouter()

	r.Run(":8080")
}
