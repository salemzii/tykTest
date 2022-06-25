package dbs

import (
	"context"
	"fmt"

	"github.com/salemzii/tykTest/files"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Client
}

var (
	tykCollection *mongo.Collection
)

//
func NewMongoRepository(db *mongo.Client) *MongoRepository {
	return &MongoRepository{
		db: db,
	}
}

// Makes migration for mongodb collection
func (repo *MongoRepository) Migrate() error {
	tykCollection = repo.db.Database("testing").Collection("tyk")
	return nil
}

// Add record to mongodb collection
func (repo *MongoRepository) Create(data *files.Data) (CreatedData *files.Data, err error) {
	defer wg.Done()

	fmt.Println("Writing to mongodb")
	bson_data := bson.D{{"api_id", data.Api_Id}, {"hits", data.Hits}}
	result, err := tykCollection.InsertOne(context.TODO(), bson_data)

	if err != nil {
		return nil, ErrCreateFailed
	}
	if result.InsertedID == 0 {
		return nil, ErrCreateFailed
	}
	return data, nil
}
