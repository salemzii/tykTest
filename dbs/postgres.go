package dbs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/salemzii/tykTest/files"
)

// Makes connection to postgres server and also makes migration
func PreparePostgres() {
	defer InitWaitgroup.Done()
	Postgres_uri := os.Getenv("PG_URI")

	Postgresdb, err = sql.Open("postgres", Postgres_uri)
	if err != nil {
		log.Fatal(errors.New(ErrConnectionFailed.Error() + ": " + err.Error()))
	}
	fmt.Println("Postgresdb is active")
	if err := MigratePostgres(Postgresdb); err != nil {
		log.Fatal(errors.New(ErrMigrationFailed.Error() + ": " + err.Error()))
	}
}

// Migrates table for storing our Data
func MigratePostgres(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tykdata(
			id SERIAL PRIMARY KEY,
			api_id varchar(10) NOT NULL,
			hits integer NOT NULL
		);
	`
	_, err := Postgresdb.Exec(query)

	return err
}

// Adds a data record to postgresdb
func AddDataRecordPostgres(db *sql.DB, data *files.Data) (CreatedData *files.Data, err error) {
	defer wg.Done()
	fmt.Println("Writing to Postgres")
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// execute insert statement
	_, err = tx.ExecContext(ctx, "INSERT INTO tykdata(api_id, hits) values($1, $2)", data.Api_Id, data.Hits)

	if err != nil {
		tx.Rollback()
		return nil, errors.New(ErrCreateFailed.Error() + ": " + err.Error())
	}

	if err = tx.Commit(); err != nil {
		return nil, errors.New(ErrCreateFailed.Error() + ": " + err.Error())
	}

	return data, nil
}
