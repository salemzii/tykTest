package dbs

import (
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/salemzii/tykTest/files"
	"github.com/stretchr/testify/assert"
)

var (
	data = files.Data{Api_Id: "ghjtWQR", Hits: 4}
)

func LoadTest() {
	if len(os.Args) > 1 && os.Args[1][:5] == "-test" {
		log.Println("testing")
		return
	}
}

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

func TestAddDataRecordMongodb(t *testing.T) {
	// initialize mock collection type for test
	mockCol := &mockCollection{}
	wg.Add(1)

	// Make call to add record to mongodb collection, but with mockCollection struct
	data, err := AddDataRecordMongodb(mockCol, &data)
	wg.Wait()
	assert.Nil(t, err)
	assert.NotNil(t, data.Api_Id)
}
