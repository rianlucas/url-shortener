package repositories

import (
	"context"
	"time"

	"github.com/rianlucas/url-shortener/internal/database"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (c *ClickAnalyticsRepository) Create(createClickDto *dto.CreateClickDto) (models.ClickAnalytics, error) {

	click := models.ClickAnalytics{
		UrlId:     createClickDto.UrlId,
		Ip:        createClickDto.Ip,
		Country:   createClickDto.Country,
		City:      createClickDto.City,
		Browser:   createClickDto.Browser,
		Os:        createClickDto.Os,
		Timezone:  createClickDto.Timezone,
		ClickedAt: time.Now(),
	}

	result, err := c.collection.InsertOne(c.ctx, click)
	if err != nil {
		return models.ClickAnalytics{}, err
	}

	if id, ok := result.InsertedID.(bson.ObjectID); ok {
		click.Id = id
	}

	return click, nil
}
