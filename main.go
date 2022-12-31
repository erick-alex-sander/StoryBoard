package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type TitleStory struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	db, _ = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&TitleStory{})

	r := gin.Default()
	r.GET("/titlestories", GetTitles)
	r.GET("/titlestories/:id", GetTitle)
	r.POST("/titlestories", CreateTitle)
	r.PUT("/titlestories/:id", UpdateTitle)
	r.DELETE("/titlestories/:id", DeleteTitle)
	r.Run(":8080")

}

func GetTitles(c *gin.Context) {
	var titleStories []TitleStory
	if err := db.Find(&titleStories).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, titleStories)
	}

}

func GetTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var titleStory TitleStory
	if err := db.Where("id = ?", id).First(&titleStory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, titleStory)
	}
}

func CreateTitle(c *gin.Context) {
	var titleStory TitleStory
	c.BindJSON(&titleStory)

	db.Create(&titleStory)
	c.JSON(200, titleStory)
}

func UpdateTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var titleStory TitleStory

	if err := db.Where("id = ?", id).First(&titleStory).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&titleStory)

	db.Save(&titleStory)
	c.JSON(200, titleStory)
}

func DeleteTitle(c *gin.Context) {
	id := c.Params.ByName("id")
	var titleStory TitleStory

	d := db.Where("id = ?", id).Delete(&titleStory)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
