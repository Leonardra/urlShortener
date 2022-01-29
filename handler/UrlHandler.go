package handler

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	"urlShortener/model"
	"urlShortener/model/dtos"
)

func GetPing(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"Hello":"Ping",
	})
}

func ShortenUrl(urlRepository *model.UrlDataBase) gin.HandlerFunc{
	urlRequest := new(dtos.UrlRequestDto)

	return func(c *gin.Context){

		err := c.Bind(urlRequest)
		urlModel := model.Url{
			Id:           urlRepository.GenerateId(),
			OriginalUrl:  urlRequest.OriginalUrl,
			ShortenedUrl: generateRandomString(),
		}
		urlRepository.SaveUrl(&urlModel)
		c.JSON(http.StatusOK, urlModel)
		if err != nil {
			return
		}
	}
}

func generateRandomString() string{
	var result =""
	rand.Seed(time.Now().UnixNano())
	letterArray := []string{"a","b","c","d","e","f","g","h" +
		"i","j","k","l","m","n","0","p","q" +
		"r","s","t","u","v","w","x","y","z","1","2","3","4","5","6" +
		"7","8","16","9" +
		"0","A","B","C","D","E","F","G","H" +
		"I","J","K","L","M","N","O","P","Q","R","S","T","U" +
		"V","W","X","Y","Z"}
	for i:= 0; i < 8; i++{
		randomNumber := rand.Intn(len(letterArray))
		result+=letterArray[randomNumber]
	}
	return result
}
