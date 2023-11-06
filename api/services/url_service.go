package services

import (
	"os"

	"github.com/andresromeroh/url-shortener/api/db"
	"github.com/andresromeroh/url-shortener/api/dto"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UrlService struct {
	*BaseService
}

func NewUrlService() *UrlService {
	baseService := NewBaseService()
	return &UrlService{BaseService: baseService}
}

func (u *UrlService) CreateUrl(createUrlDto dto.CreateUrl) (*db.URLModel, error) {
	nanoid, _ := gonanoid.New()
	shortUrl := os.Getenv("BASE_URL") + nanoid

	createdUrl, err := u.PrismaClient.URL.CreateOne(
		db.URL.ShortURL.Set(shortUrl),
		db.URL.LongURL.Set(createUrlDto.LongURL),
	).Exec(u.ctx)

	if err != nil {
		return nil, err
	}

	return createdUrl, nil
}
