package api

import (
	"example/web-service-gin/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func AddAlbums(c *gin.Context) {
	var newAlbum album
	c.BindJSON(&newAlbum)
	fmt.Println(newAlbum)
	q := fmt.Sprintf("INSERT INTO public.albums(id, title, artist, price) VALUES (%d, '%s', '%s', %f);", 3, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	_, err := database.Db.Query(q)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Couldn't create the new album.")
	} else {
		c.JSON(http.StatusOK, "User is successfully created.")
	}
}
