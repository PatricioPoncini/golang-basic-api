package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Sour", Artist: "Olivia Rodrigo", Price: 5.99},
	{ID: "2", Title: "Fuerza Natural", Artist: "Gustavo Cerati", Price: 5.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}
