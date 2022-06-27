package dbs

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/salemzii/tykTest/files"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Accepts any type that implements method InsertOne()
type CollectionApi interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

// Fake Collection struct that implements CollectionApi interface
type mockCollection struct {
}

// Fake method to simulate *mongo.Collection.InsertOne method
func (m *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	collection := &mongo.InsertOneResult{}

	return collection, nil
}

var (
	ErrConnectionFailed = errors.New("error connecting to db: ")
	ErrDuplicate        = errors.New("record already exists")
	ErrCreateFailed     = errors.New("failed to create record: ")
	ErrMigrationFailed  = errors.New("failed to migrate db: ")

	Postgresdb *sql.DB
	client     *mongo.Client
	err        error

	wg            sync.WaitGroup
	InitWaitgroup sync.WaitGroup
)

// asynchronously initialize databases connections
func init() {
	if len(os.Args) > 1 && os.Args[1][:5] == "-test" {
		log.Println("testing")
		return
	} else {
		InitWaitgroup.Add(2)
		go PreparePostgres()
		go PrepareMongo()
		InitWaitgroup.Wait()
	}
}

// this function, recieves a list of Data type,
// iterates over them and concurrently passes each data to
// respective databases i.e(mongodb and postgresql)

func WriteData(data []files.Data) {

	wg.Add(len(data))

	for _, v := range data {
		go AddDataRecordPostgres(Postgresdb, &v)
		go AddDataRecordMongodb(tykCollection, &v)
	}

	wg.Wait()
}
