package dbs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/salemzii/tykTest/files"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {

	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS TykData(
			id SERIAL PRIMARY KEY,
			api_id varchar(10) NOT NULL,
			hits integer NOT NULL
		);
	`
	_, err := repo.db.Exec(query)

	return err
}

func (repo *PostgresRepository) Create(data *files.Data) (CreatedData *files.Data, err error) {
	defer wg.Done()
	fmt.Println("Writing to Postgres")

	stmt, err := repo.db.Prepare("INSERT INTO TykData(api_id, hits) values($1, $2)")
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
