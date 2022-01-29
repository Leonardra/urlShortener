package main

import (
	"github.com/gin-gonic/gin"
	"urlShortener/handler"
	"urlShortener/model"
)
func main() {
	r := gin.Default()

	db := model.CreateUrlDataBase()
	r.GET("/", handler.GetPing)
	r.POST("/", handler.ShortenUrl(db))
	err := r.Run()
	if err != nil {
		return 
	}
}
