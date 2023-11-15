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
		db.URL.ID.Set(nanoid),
		db.URL.ShortURL.Set(shortUrl),
		db.URL.LongURL.Set(createUrlDto.LongURL),
		db.URL.UserIP.Set(createUrlDto.IP),
	).Exec(u.ctx)

	if err != nil {
		return nil, err
	}

	return createdUrl, nil
}

func (u *UrlService) FindById(id string) (*db.URLModel, error) {
	url, err := u.PrismaClient.URL.FindUnique(
		db.URL.ID.Equals(id),
	).Exec(u.ctx)

	if err != nil {
		return nil, err
	}

	return url, nil
}
