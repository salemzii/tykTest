package dbs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/salemzii/tykTest/files"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	tykCollection *mongo.Collection
)

// connects to mongodb server and defines value for Collection
func PrepareMongo() {
	defer InitWaitgroup.Done()
	mongo_uri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to mongodb client
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(errors.New(ErrConnectionFailed.Error() + ": " + err.Error()))
	}
	fmt.Println("mongodb is active")
	if err := MigrateMongodb(client); err != nil {
		log.Fatal(errors.New(ErrMigrationFailed.Error() + ": " + err.Error()))
	}
}

// Makes migration for mongodb collection
func MigrateMongodb(db *mongo.Client) error {
	tykCollection = db.Database("testing").Collection("tyk")
	return nil
}

// Adds record to a mongodb collection
// implements CollectionApi type
func AddDataRecordMongodb(collection CollectionApi, data *files.Data) (CreatedData *files.Data, err error) {
	defer wg.Done()

	fmt.Println("Writing to mongodb")

	bson_data := bson.D{{Key: "api_id", Value: data.Api_Id}, {Key: "hits", Value: data.Hits}}
	result, err := collection.InsertOne(context.TODO(), bson_data)

	if err != nil {
		return nil, ErrCreateFailed
	}
	if result.InsertedID == 0 {
		return nil, ErrCreateFailed
	}
	return data, nil
}
