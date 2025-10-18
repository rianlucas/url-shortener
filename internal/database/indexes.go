package database

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateUrlIndexes(db *mongo.Database) error {
	collection := db.Collection(UrlCollection)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "shortCode", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}
