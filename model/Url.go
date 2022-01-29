package model

type Url struct{
	Id          int   `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
}


type UrlDataBase struct{
	idGenerated int
	urls map[int]Url
}

func createUrlDataBase() *UrlDataBase{
  urlDataBase := new(UrlDataBase)
  urlDataBase.urls = make(map[int]Url)
  urlDataBase.idGenerated = 0
  return urlDataBase
}

func (urlDb *UrlDataBase) generateId() int{
	urlDb.idGenerated = urlDb.idGenerated + 1
	return urlDb.idGenerated
}

func (urlDb *UrlDataBase) saveUrl(url *Url){
	url.Id = urlDb.generateId()
	urlDb.urls[url.Id] = *url
}

func (urlDb *UrlDataBase) getAll() map[int]Url{
	return urlDb.urls
}
