package _interface

import (
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepositoryInterface interface {
	Create(dto dto.CreateUrlDto) *mongo.InsertOneResult
	Update(url models.Url) (bool, error)
	FindByShortCode(shortCode string) (models.Url, error)
}
