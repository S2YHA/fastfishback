package main

import (
    "fastfishback/config"
    "fastfishback/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Initialize the database
    config.ConnectDatabase()

    // Register routes
    routes.RegisterUserRoutes(router)

    // Start the server
    router.Run(":8080")
}
