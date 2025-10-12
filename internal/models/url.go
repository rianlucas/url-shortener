package models

import "time"

type Url struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	LongUrl     string    `bson:"longUrl" json:"longUrl"`
	ShortCode   string    `bson:"shortCode" json:"shortCode"`
	AccessCount int       `bson:"accessCount" json:"accessCount"`
	CreatedAt   time.Time `bson:"createdAt" json:"-"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"-"`
}
