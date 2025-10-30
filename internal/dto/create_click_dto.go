package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateClickDto struct {
	UrlId   bson.ObjectID `bson:"urlId" json:"urlId"`
	Ip      string        `bson:"ip" json:"ip"`
	Country string        `bson:"country" json:"country"`
	City    string        `bson:"city" json:"city"`
	Device  string        `bson:"device" json:"device"`
	Browser string        `bson:"browser" json:"browser"`
	Os      string        `bson:"os" json:"os"`
	Referer string        `bson:"referer" json:"referer"`
}
