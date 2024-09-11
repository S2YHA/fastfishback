package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"os"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type User struct {
	Id   int
	Name string
}

type Lesson struct {
	UserId int
	Name   string
}

type Words struct {
	Id       int
	LessonId string
	WordPl   string
	WordEn   string
}

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	// Define MySQL connection string
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	fmt.Println(dsn)

	// Open the MySQL connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Connected to MySQL database!")
}
