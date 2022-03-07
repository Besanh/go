package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gallery struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
}

var gal = []Gallery{
	{ID: "1", Name: "gallery 1", Size: "13kB"},
	{ID: "2", Name: "gallery 2", Size: "1MB"},
	{ID: "3", Name: "gallery 3", Size: "0kB"},
}

func main() {
	route := gin.Default()
	route.GET("/gallery", getGallery)
	route.POST("/gallery", postGallery)
	route.GET("/gallery/:id", getGalleryByID)

	route.Run("localhost:8080")
}

func getGallery(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, gal)
}

func postGallery(a *gin.Context) {
	var postGal Gallery
	if err := a.BindJSON(&postGal); err != nil {
		return
	}

	gal = append(gal, postGal)
	a.IndentedJSON(http.StatusCreated, postGal)
}

func getGalleryByID(a *gin.Context) {
	id := a.Param("id")

	for _, s := range gal {
		if s.ID == id {
			a.IndentedJSON(http.StatusOK, s)
			return
		}
	}
	a.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}
