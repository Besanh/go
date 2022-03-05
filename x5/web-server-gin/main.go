package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "A1", Artist: "Anh Le", Price: 10},
	{ID: "2", Title: "A2", Artist: "Huyen Thu", Price: 20.1},
	{ID: "3", Title: "A3", Artist: "Tuyet", Price: 30.5},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/get-album-by-id/:id", getAlbumById)

	router.Run("localhost:8080")
}

// Get
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Post
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Them albums vao newAlbum
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Get by id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "data not found"})
}
