package main

import (
	"example/web-service-gin/api"
	"example/web-service-gin/database"
	"example/web-service-gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// var albums = []album{
// 	{ID: Helper.generateUID(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: Helper.generateUID(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: Helper.generateUID(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.AddAlbums)
	router.GET("/albums/:id", getAlbumByID)

	database.ConnectDatabase()
	router.Run("localhost:8080")
}

// func postAlbum(c *gin.Context) {
// 	var newAlbum album
// 	data, err := c.GetRawData()

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	newAlbum.ID = Helper.generateUID()
// 	albums = append(albums, newAlbum)

// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

func getAlbumByID(c *gin.Context) {
	// id := c.Param("id")

	// for _, a := range albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
