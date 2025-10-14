package service

import (
	"fmt"

	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"github.com/rianlucas/url-shortener/pkg/shortcode"
)

type UrlService struct {
	repository *repositories.UrlRepository
}

func NewUrlService(repository *repositories.UrlRepository) *UrlService {
	return &UrlService{repository: repository}
}

func (u *UrlService) Create(urlDto dto.CreateUrlDto) (models.Url, error) {

	generatedCode, err := shortcode.Generate(6)
	if err != nil {
		return models.Url{}, fmt.Errorf("error generating short code: %w", err)
	}
	urlDto.ShortCode = generatedCode

	return u.repository.Create(urlDto)
}
