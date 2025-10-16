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

func (u *UrlService) FindByShortCode(shortCode string) (models.Url, error) {
	result, err := u.repository.FindByShortCode(shortCode)
	if err != nil {
		return models.Url{}, err
	}

	return result, nil
}

func (u *UrlService) Update(url models.Url) (bool, error) {
	return u.repository.Update(url)
}
