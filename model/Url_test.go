package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_testThatUrlCanBeAdded(t *testing.T){
	urlDb := createUrlDataBase()
	url := new(Url)
	url.OriginalUrl = "www.testing.com/aString"
	url.ShortenedUrl = "www.testing.com/y6Ghty"
	urlDb.saveUrl(url)
	assert.Equal(t,1, len(urlDb.urls))
	assert.Equal(t, 1, url.Id)
}
