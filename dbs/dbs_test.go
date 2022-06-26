package dbs

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/salemzii/tykTest/files"
)

var (
	data = files.Data{Api_Id: "ghjtWQR", Hits: 4}
)

// a successful case
func TestAddDataRecordPostgres(t *testing.T) {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockdb.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO tykdata(api_id, hits) values($1, $2)")).WithArgs(data.Api_Id, data.Hits).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	wg.Add(1)

	if _, err = AddDataRecordPostgres(mockdb, &data); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	wg.Wait()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
