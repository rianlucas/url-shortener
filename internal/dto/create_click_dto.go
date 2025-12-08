package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateClickDto struct {
	UrlId    bson.ObjectID `bson:"urlId"`
	Ip       string        `bson:"ip" json:"ip"`
	Country  string        `bson:"country" json:"country"`
	City     string        `bson:"city" json:"city"`
	Browser  string        `bson:"browser" json:"browser"`
	Os       string        `bson:"os" json:"os"`
	Timezone string        `bson:"timezone" json:"timezone"`
}
