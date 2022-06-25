package dbs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/salemzii/tykTest/files"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Migrate() error
	Create(data files.Data) (CreatedData *files.Data, err error)
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrCreateFailed = errors.New("failed to create record")

	Postgresdb *sql.DB
	client     *mongo.Client
	err        error

	wg            sync.WaitGroup
	InitWaitgroup sync.WaitGroup
)

func init() {
	InitWaitgroup.Add(2)
	go PreparePostgres()
	go PrepareMongo()
	InitWaitgroup.Wait()
}

func PreparePostgres() {
	defer InitWaitgroup.Done()
	Postgres_uri := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", "postgresql", "postgres", "auth1234", "localhost:5432", "contact")

	Postgresdb, err = sql.Open("postgres", Postgres_uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Postgresdb is active")
	if err := MigratePostgres(Postgresdb); err != nil {
		log.Fatal(err)
	}
}

func PrepareMongo() {
	defer InitWaitgroup.Done()
	mongo_uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb is active")
	if err := MigrateMongodb(client); err != nil {
		log.Fatal(err)
	}
}

func WriteData(data []files.Data) {

	wg.Add(len(data))

	for _, v := range data {
		go AddDataRecordPostgres(Postgresdb, &v)
		go AddDataRecordMongodb(&v)
	}

	wg.Wait()
	fmt.Println("Finished writing to mongodb and postgres")

}
