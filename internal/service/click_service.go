package service

import (
	"context"

	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
)

type ClickService struct {
	ctx        context.Context
	repository repositories.ClickAnalyticsRepository
}

func (c *ClickService) Create(clickDto dto.CreateClickDto) (models.ClickAnalytics, error) {
	result, err := c.repository.Create(clickDto)
	if err != nil {
		return models.ClickAnalytics{}, err
	}

	return result, nil
}
