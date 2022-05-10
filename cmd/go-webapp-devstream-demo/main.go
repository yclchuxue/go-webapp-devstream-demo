package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yclchuxue/go-webapp-devstream-demo/internal/pkg/album"
)

func main() {
	router := gin.Default()

	router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)

	router.Run("localhost:8080")
}
