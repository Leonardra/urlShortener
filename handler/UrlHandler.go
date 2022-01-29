package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPing(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"Hello":"Ping",
	})
}

func shortenUrl(urlString string){

}
