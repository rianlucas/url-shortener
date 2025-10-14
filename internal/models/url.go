package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Url struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LongUrl     string        `bson:"longUrl" json:"longUrl"`
	ShortCode   string        `bson:"shortCode" json:"shortCode"`
	AccessCount int           `bson:"accessCount" json:"accessCount"`
	CreatedAt   time.Time     `bson:"createdAt" json:"-"`
	UpdatedAt   time.Time     `bson:"updatedAt" json:"-"`
}
