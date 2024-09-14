package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/joho/godotenv"
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

var db *gorm.DB

func main() {
	connectToDb()

	router := gin.Default()

	// Definiowanie prostego endpointu
	router.GET("/users", getUsers)

	// Uruchomienie serwera na porcie 8080
	router.Run(":8080")
}

func connectToDb() *gorm.DB {
	var err error

	err = godotenv.Load()
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

	// Open the MySQL connection
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	fmt.Println("Connected to MySQL database!")

	return db
}

func getUsers(c *gin.Context) {
	var users []User
	// Query the database to get all records from the "users" table
	result := db.Find(&users)

	// Check if any error occurred during the query
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with the list of users as JSON
	c.JSON(http.StatusOK, users)
}
