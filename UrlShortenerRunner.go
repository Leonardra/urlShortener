package main

import (
	"github.com/gin-gonic/gin"
	"urlShortener/handler"
)
func main() {
	r := gin.Default()
	r.GET("/", handler.GetPing)
	err := r.Run()
	if err != nil {
		return 
	}
}
