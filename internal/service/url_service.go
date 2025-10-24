package service

import (
	"errors"
	"fmt"

	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"github.com/rianlucas/url-shortener/pkg/shortcode"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlService struct {
	repository *repositories.UrlRepository
}

func NewUrlService(repository *repositories.UrlRepository) *UrlService {
	return &UrlService{repository: repository}
}

func (u *UrlService) GenerateQrCode(url string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}

func (u *UrlService) Create(urlDto dto.CreateUrlDto) (models.Url, error) {

	generatedCode, err := shortcode.Generate(6)
	if err != nil {
		return models.Url{}, fmt.Errorf("error generating short code: %w", err)
	}

	maxAttempts := 3
	for i := 1; i < maxAttempts; i++ {
		fmt.Printf("tentativa: %v\n shortCode: %s\n", i, generatedCode)
		result, err := u.FindByShortCode(generatedCode)
		if errors.Is(err, mongo.ErrNoDocuments) {
			break
		}
		if result.ShortCode == generatedCode {
			generatedCode, _ = shortcode.Generate(6)
		}
		if i == maxAttempts-1 && result.ShortCode == generatedCode {
			return models.Url{}, fmt.Errorf("error generating short code")
		}
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
