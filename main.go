package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//config router webservice
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8000")
}

//////////////HTTP VERBS//////////////////

//return a JSON data
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

//return a specific item of album model
func getAlbumById(context *gin.Context) {
	id := context.Param("id")

	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found, try again..."})
}

//create album item and append to albums array
func postAlbum(context *gin.Context) {
	var newAlbum album

	if error := context.BindJSON(&newAlbum); error != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusOK, newAlbum)
}

//model structure
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//seed data of album model
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
