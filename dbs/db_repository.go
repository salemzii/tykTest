package dbs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/salemzii/tykTest/files"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Migrate() error
	Create(data files.Data) (CreatedData *files.Data, err error)
}

var (
	ErrDuplicate            = errors.New("record already exists")
	ErrCreateFailed         = errors.New("failed to create record")
	writePostgresRepository *PostgresRepository
	writeMongoRepository    *MongoRepository

	wg sync.WaitGroup
)

func init() {
	go func() {
		Postgres_uri := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", "postgresql", "postgres", "auth1234", "localhost:5432", "contact")

		db, err := sql.Open("postgres", Postgres_uri)
		if err != nil {
			log.Fatal(err)
		}

		writePostgresRepository = NewPostgresRepository(db)
		if err := writePostgresRepository.Migrate(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		mongo_uri := "localhost:"
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
		if err != nil {
			log.Fatal(err)
		}

		writeMongoRepository = NewMongoRepository(client)
	}()

}

func WriteData(data *files.Data) {
	wg.Add(2)
	go writePostgresRepository.Create(data)
	go writeMongoRepository.Create(data)
	wg.Wait()

	fmt.Println("Finished writing to mongo and postgres")
}
