package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rianlucas/url-shortener/internal/database"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClickAnalyticsRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewClickAnalyticsRepository(ctx context.Context, db *mongo.Database) *ClickAnalyticsRepository {
	return &ClickAnalyticsRepository{
		ctx:        ctx,
		collection: db.Collection(database.ClickAnalyticsCollection),
	}
}

func (c *ClickAnalyticsRepository) Create(createClickDto dto.CreateClickDto) {
	var clickAnalytics models.ClickAnalytics

}
