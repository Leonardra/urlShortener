package model

type Url struct{
	Id          int   `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
}

