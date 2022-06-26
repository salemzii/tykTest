package dbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
