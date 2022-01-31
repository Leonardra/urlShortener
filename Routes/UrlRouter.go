package Routes

import (
	"github.com/gin-gonic/gin"
	"urlShortener/controllers"
	"urlShortener/model"
)

func UrlRouters(router *gin.Engine){
	db := model.CreateUrlDataBase()
	router.GET("/", controllers.GetPing)
	//r.GET("/:id", controllers.GetUrl(db))
	router.POST("/", controllers.ShortenUrl(db))
}