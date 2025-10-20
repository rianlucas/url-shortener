package database

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateUrlIndexes(db *mongo.Database) error {
	collection := db.Collection(UrlCollection)

	// Unique index on shortCode
	uniqueIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "shortCode", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), uniqueIndexModel)
	if err != nil {
		return err
	}

	ttlIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiresAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	_, err = collection.Indexes().CreateOne(context.Background(), ttlIndexModel)
	return err
}
