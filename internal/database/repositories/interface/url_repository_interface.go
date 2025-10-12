package _interface

import (
	"github.com/rianlucas/url-shortener/internal/dto"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepositoryInterface interface {
	Create(dto dto.CreateUrlDto) *mongo.InsertOneResult
	Update() string
}
