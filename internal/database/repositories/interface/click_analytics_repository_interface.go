package _interface

import (
	"github.com/rianlucas/url-shortener/internal/dto"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClickAnalyticsRepositoryInterface interface {
	Create(createClickDto dto.CreateClickDto) *mongo.InsertOneResult
}
