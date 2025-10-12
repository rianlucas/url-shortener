package service

import (
	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
)

type UrlService struct {
	repository *repositories.UrlRepository
}

func CreateUrlService(repository *repositories.UrlRepository) *UrlService {
	return &UrlService{repository: repository}
}

func (u *UrlService) Create(urlDto dto.CreateUrlDto) models.Url {
	return u.repository.Create(urlDto)
}
