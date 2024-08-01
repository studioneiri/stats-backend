package database

import (
   "fmt"
   "os"
   "strconv"
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    db *gorm.DB
)

// Uppercase allows func to be called outside of the directory
func Connect() {
    var err error
    host := os.Getenv("DB_HOST")
    port, _ := strconv.Atoi(os.Getenv("DB_PORT")) // convert to int since port is int type
    user := os.Getenv("DB_USER")
    dbname := os.Getenv("DB_NAME")
    pass := os.Getenv("DB_PASSWORD")

    // TODO: Enable SSL verification
    dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, dbname)
    db, err = gorm.Open("postgres", dsn)
    if err != nil {
        panic("Failed to connect to database")
    } else {
        fmt.Println("Successfully connected to database!")
    }
}

type Location struct {
    ID  string `json:"id"`
    Name string `json:"name"`
    Latitude float32 `sql:"type:decimal(2,15);" json:"latitude"`
    Longitude float32 `sql:"type:decimal(3,15);" json:"longitude"`
    PeakPopulation int `json:"peakPopulation"`
    PeakYear int `json:"peakYear"`
}

type Language struct {
    ID  string `json:"id"`
    Language string `json:"language"`
    EstimatedSpeakers int `json:"estimatedSpeakers"`
}

func GetLocations(c *gin.Context) {
    var locations []Location
    db.Find(&locations)
    c.JSON(200, locations)
}

func GetLanguages(c *gin.Context) {
    var languages []Language
    db.Find(&languages)
    c.JSON(200, languages)
}

func CreateLocation(c *gin.Context) {
    var location Location
    if err := c.BindJSON(&location); err != nil {
        c.AbortWithStatus(400)
        return
    }
    db.Create(&location)
    c.JSON(201, location)
}

func DeleteLocationByID(c *gin.Context) {
    id := c.Param("id")
    var location Location
    db.Delete(&location, "id = ?", id)
    c.Status(204)
}

func GetLocationByID(c *gin.Context) {
    id := c.Param("id")
    var location Location
    db.First(&location, "id = ?", id)
    c.JSON(200, location)
}

func GetLanguagesById(c *gin.Context) {
    id := c.Param("id")
    var languages []Language
    db.Find(&languages, "id = ?", id)
    c.JSON(200, languages)
}