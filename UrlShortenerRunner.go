package main

import (
	"github.com/gin-gonic/gin"
	"urlShortener/Routes"
)
func main() {
	router := gin.Default()

	Routes.UrlRouters(router)
	err := router.Run()
	if err != nil {
		return 
	}
}
