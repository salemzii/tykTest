package dbs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/salemzii/tykTest/files"
)

func MigratePostgres(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tykdata(
			id SERIAL PRIMARY KEY,
			api_id varchar(10) NOT NULL,
			hits integer NOT NULL
		);
	`
	_, err := db.Exec(query)

	return err
}

func AddDataRecordPostgres(db *sql.DB, data *files.Data) (CreatedData *files.Data, err error) {
	defer wg.Done()
	fmt.Println("Writing to Postgres")

	stmt, err := db.Prepare("INSERT INTO tykdata(api_id, hits) values($1, $2)")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := stmt.Exec(data.Api_Id, data.Hits)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	rowsaffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsaffected == 0 {
		return nil, ErrCreateFailed
	}

	return data, nil
}
