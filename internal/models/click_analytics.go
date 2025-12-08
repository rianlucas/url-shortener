package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ClickAnalytics struct {
	Id        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UrlId     bson.ObjectID `bson:"urlId" json:"urlId"`
	Ip        string        `bson:"ip" json:"ip"`
	Country   string        `bson:"country" json:"country"`
	City      string        `bson:"city" json:"city"`
	Browser   string        `bson:"browser" json:"browser"`
	Os        string        `bson:"os" json:"os"`
	Timezone  string        `bson:"timezone" json:"timezone"`
	ClickedAt time.Time     `bson:"clickedAt" json:"clickedAt"`
}
