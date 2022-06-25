package dbs

import (
	"context"
	"fmt"

	"github.com/salemzii/tykTest/files"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	tykCollection *mongo.Collection
)

// Makes migration for mongodb collection
func MigrateMongodb(db *mongo.Client) error {
	tykCollection = db.Database("testing").Collection("tyk")
	return nil
}

// Adds record to mongodb collection
func AddDataRecordMongodb(data *files.Data) (CreatedData *files.Data, err error) {
	defer wg.Done()

	fmt.Println("Writing to mongodb")

	bson_data := bson.D{{Key: "api_id", Value: data.Api_Id}, {Key: "hits", Value: data.Hits}}
	result, err := tykCollection.InsertOne(context.TODO(), bson_data)

	if err != nil {
		return nil, ErrCreateFailed
	}
	if result.InsertedID == 0 {
		return nil, ErrCreateFailed
	}
	return data, nil
}
