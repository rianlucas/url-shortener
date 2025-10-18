package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/rianlucas/url-shortener/internal/database"
	"time"

	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewUrlRepository(ctx context.Context, db *mongo.Database) *UrlRepository {
	return &UrlRepository{
		ctx:        ctx,
		collection: db.Collection(database.UrlCollection),
	}
}

func (u *UrlRepository) Create(urlDto dto.CreateUrlDto) (models.Url, error) {

	newUrl := models.Url{
		LongUrl:     urlDto.LongUrl,
		ShortCode:   urlDto.ShortCode,
		AccessCount: urlDto.AccessCount,
	}

	now := time.Now()
	newUrl.CreatedAt = now
	newUrl.UpdatedAt = now

	fmt.Printf("Creating new URL: %+v\n", newUrl)

	result, err := u.collection.InsertOne(u.ctx, newUrl)
	if err != nil {
		return models.Url{}, fmt.Errorf("failed to insert URL into database: %w", err)
	}

	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		newUrl.ID = oid
	}

	return newUrl, nil
}

func (u *UrlRepository) Update(url models.Url) (bool, error) {
	update := bson.D{
		{
			"$inc", bson.D{
				{"accessCount", 1},
			},
		},
	}

	result, err := u.collection.UpdateByID(u.ctx, url.ID, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (u *UrlRepository) FindByShortCode(shortCode string) (models.Url, error) {
	var urlEntity models.Url

	result := u.collection.FindOne(u.ctx, bson.M{"shortCode": shortCode})

	err := result.Decode(&urlEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Url{}, fmt.Errorf("URL with shortCode '%s' not found: %w", shortCode, err)
		}
		return models.Url{}, fmt.Errorf("failed to decode URL: %w", err)
	}

	return urlEntity, nil
}
