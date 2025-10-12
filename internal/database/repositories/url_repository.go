package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UrlRepository struct {
	ctx    context.Context
	Client *mongo.Client
}

func CreateUrlRepository(ctx context.Context, client *mongo.Client) *UrlRepository {
	return &UrlRepository{
		ctx:    ctx,
		Client: client,
	}
}

func (s *UrlRepository) Create(urlDto dto.CreateUrlDto) models.Url {
	collection := s.Client.Database("url-shortener").Collection("urls")

	newUrl := models.Url{
		LongUrl:     urlDto.LongUrl,
		ShortCode:   urlDto.ShortCode,
		AccessCount: urlDto.AccessCode,
	}

	now := time.Now()
	newUrl.CreatedAt = now
	newUrl.UpdatedAt = now

	result, err := collection.InsertOne(s.ctx, newUrl)
	if err != nil {
		fmt.Println(err)
	}

	id := fmt.Sprintf("%v", result.InsertedID)
	id = strings.TrimPrefix(id, "ObjectID(\"")
	id = strings.TrimSuffix(id, "\")")

	newUrl.ID = id
	return newUrl
}

func (s *UrlRepository) Update() string {
	//TODO implement me
	panic("implement me")
}
