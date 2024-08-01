package main

import (
    "github.com/gin-gonic/gin"
    "studio.neiri/database"
)

func init() {
    // Connect to database
    database.Connect()

}

func main() {
    r := gin.Default()
    // Routes
    r.GET("/locations", database.GetLocations)
    r.GET("/languages", database.GetLanguages)
    r.POST("/locations", database.CreateLocation)
    r.GET("/locations/:id", database.GetLocationByID)
    r.GET("/languages/:id", database.GetLanguagesById)
    r.DELETE("/locations/:id", database.DeleteLocationByID)

    // Start server
    r.Run("localhost:8080")
}
