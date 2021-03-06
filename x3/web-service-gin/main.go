package main

import (
	// Gin FW
	"github.com/gin-gonic/gin"
	// Http web
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Tra ve 1 instance(lop con) cua Engine cung voi middleware cua Logger va Recovery
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	// Dinh kem router toi http.Server va start srv
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Dung de serialize(tuan tu hoa) struct vao JSON va them chung vao response
	c.IndentedJSON(http.StatusOK, albums)
}

/**
* Them album vao response
* Khi restart lai srv thi moi hien
 */
func postAlbums(c *gin.Context) {
	// Khai bao newAlbum kieu struct
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	// bind request body vao newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum) // status 201
}

// Lay album theo id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		// ID nay tuong ung voi khai bao kieu struct(khong phai id)
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
