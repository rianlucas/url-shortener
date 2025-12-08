package database

import "github.com/rianlucas/url-shortener/internal/models"

type ClickBuilder struct {
	ClickAnalytics *models.ClickAnalytics
}

func NewClickBuilder(model *models.ClickAnalytics) *ClickBuilder {
	return &ClickBuilder{
		ClickAnalytics: model,
	}
}
