package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codeshaine/url-shortner/db"
	"github.com/codeshaine/url-shortner/internal/router"
	utils "github.com/codeshaine/url-shortner/internal/utils"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("URL Shortner project Starting...")

	//loading .evn file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// port and address
	port := utils.GetIntEnv("PORT")
	addr := fmt.Sprintf("localhost:%d", port)

	//database connection
	db.DbConnect()
	// defer log.Println("disconnecting db...")
	// defer Db.Close()
	defer func() {
		if err := db.Db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}()

	//router
	mux := router.Router()

	//starting application
	log.Printf("Application running on http://%s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal("Error occured during server startup", err)
	}

}
