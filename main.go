package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "studio.neiri/database"
)

func init() {
    // Connect to database
    database.Connect()

}

func main() {
    router := gin.Default()
    router.Use(cors.Default())
    // Routes
    router.GET("/locations", database.GetLocations)
    router.GET("/languages", database.GetLanguages)
    router.POST("/locations", database.CreateLocation)
    router.GET("/locations/:id", database.GetLocationByID)
    router.GET("/languages/:id", database.GetLanguagesById)
    router.DELETE("/locations/:id", database.DeleteLocationByID)

    // Start server
    router.Run("localhost:8080")
}
