package db

import (
	"database/sql"
	"fmt"
	"log"

	utils "github.com/codeshaine/url-shortner/internal/utils"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func DbConnect() {
	var dbHost = utils.GetEnv("HOST")
	var dbPort = utils.GetIntEnv("DB_PORT")
	var dbUser = utils.GetEnv("DB_USER")
	var dbPassword = utils.GetEnv("DB_PASSWORD")
	var dbName = utils.GetEnv("DB_NAME")

	// **************** connecting to postgres *********************************
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s  sslmode=disable", dbHost, dbPort, dbUser, dbPassword)
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Databse connection Error:", err)
	}

	log.Printf("Connected to postgres db :")

	//**************** Check if the database exists and creates it if it doesnt **************************
	var exists bool
	err = Db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		Db.Close()
		log.Fatal("Error checking if database exists:", err)
	}

	if !exists {
		_, err = Db.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			Db.Close()
			log.Fatal("Error creating database:", err)
		}
		log.Printf("Database %s created successfully!", dbName)
	} else {
		log.Printf("Database %s already exists.", dbName)
	}

	// **************** connecting to postgres *********************************
	var reConErr error
	reConStr := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	Db, reConErr = sql.Open("postgres", reConStr)
	if reConErr != nil {
		log.Fatal("Databse connection Error:", err)
	}

	log.Printf("Reconnecting to : %v", dbName)

	intlErr := initialize()
	if intlErr != nil {
		log.Fatalf("Error occured while initialization of db %v", intlErr)
	}

}

func initialize() error {

	//************************** creating table initailly  **************************
	tableQuery := `
CREATE TABLE IF NOT EXISTS urls(
id SERIAL PRIMARY KEY,
long_url TEXT NOT NULL UNIQUE,
short_url TEXT NOT NULL UNIQUE,
click INTEGER DEFAULT(0),
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

	_, tbErr := Db.Exec(tableQuery)
	if tbErr != nil {
		log.Fatalf("Error occured during table creation: %v", tbErr)
	}
	log.Println("Created table if it not exists")

	//************************** creating index for the short_url field **************************
	indexQuery := `CREATE INDEX IF NOT EXISTS idx_short_url ON urls(short_url)`
	_, idxErr := Db.Exec(indexQuery)
	if idxErr != nil {
		log.Fatalf("Error occured during creation of index: %v", idxErr)
	}
	log.Println("Created index if it not exists")

	return nil
}
