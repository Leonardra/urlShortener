package model

type UrlDataBase struct{
	idGenerated int
	urls map[int]Url
}

func CreateUrlDataBase() *UrlDataBase{
	urlDataBase := new(UrlDataBase)
	urlDataBase.urls = make(map[int]Url)
	urlDataBase.idGenerated = 0
	return urlDataBase
}

func (urlDb *UrlDataBase) GenerateId() int{
	urlDb.idGenerated = urlDb.idGenerated + 1
	return urlDb.idGenerated
}

func (urlDb *UrlDataBase) SaveUrl(url *Url){
	url.Id = urlDb.GenerateId()
	urlDb.urls[url.Id] = *url
}

func (urlDb *UrlDataBase) IsUrlExist(urlString string) bool{
	for _, value := range urlDb.urls {
		if value.OriginalUrl == urlString{
			return true
		}
	}
			return false
}

func (urlDb *UrlDataBase) GetAll() map[int]Url{
	return urlDb.urls
}
