package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepository struct {
	ctx    context.Context
	Client *mongo.Client
}

func NewUrlRepository(ctx context.Context, client *mongo.Client) *UrlRepository {
	return &UrlRepository{
		ctx:    ctx,
		Client: client,
	}
}

func (u *UrlRepository) Create(urlDto dto.CreateUrlDto) (models.Url, error) {
	collection := u.Client.Database("url-shortener").Collection("urls")

	newUrl := models.Url{
		LongUrl:     urlDto.LongUrl,
		ShortCode:   urlDto.ShortCode,
		AccessCount: urlDto.AccessCount,
	}

	now := time.Now()
	newUrl.CreatedAt = now
	newUrl.UpdatedAt = now

	result, err := collection.InsertOne(u.ctx, newUrl)
	if err != nil {
		return models.Url{}, fmt.Errorf("failed to insert URL into database: %w", err)
	}

	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		newUrl.ID = oid
	}

	return newUrl, nil
}

func (u *UrlRepository) Update(url models.Url) (bool, error) {
	collection := u.Client.Database("url-shortener").Collection("urls")

	update := bson.D{
		{
			"$inc", bson.D{
				{"accessCount", 1},
			},
		},
	}

	result, err := collection.UpdateByID(u.ctx, url.ID, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (u *UrlRepository) FindByShortCode(shortCode string) (models.Url, error) {
	var urlEntity models.Url
	collection := u.Client.Database("url-shortener").Collection("urls")

	result := collection.FindOne(u.ctx, bson.M{"shortCode": shortCode})

	err := result.Decode(&urlEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Url{}, fmt.Errorf("URL with shortCode '%s' not found", shortCode)
		}
		return models.Url{}, fmt.Errorf("failed to decode URL: %w", err)
	}

	return urlEntity, nil
}
