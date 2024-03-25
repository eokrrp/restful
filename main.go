package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"restful/db"
	"strconv"
)

type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()

	database := db.Connect()
	defer database.Close()

	err := createSchema(database)
	if err != nil {
		panic(err)
	}

	router.GET("/albums", func(c *gin.Context) {
		var albums []album
		err := database.Model(&albums).Select()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, albums)
	})

	router.GET("/albums/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		album := &album{
			ID: id,
		}

		err := database.Model(album).WherePK().Select()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, album)
	})

	router.POST("/albums", func(c *gin.Context) {
		newAlbum := &album{}

		if err := c.BindJSON(newAlbum); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err.Error())
			return
		}

		if _, err := database.Model(newAlbum).Insert(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err.Error())
			return
		}

		c.JSON(http.StatusCreated, newAlbum)
	})

	err = router.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*album)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
