package handlers

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestPutEvent(t *testing.T) {
	// mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO domains \(name, events, bounced\) VALUES \(\?, \?, \?\) ON CONFLICT \(name\) DO UPDATE SET events \= domains\.events\+1, bounced \= \(CASE WHEN domains\.bounced \= TRUE THEN TRUE ELSE EXCLUDED\.bounced END\);`).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.Begin()

	// assertions
	err = putEvent(sqlxDB, "example3.com", false)
	if err != nil {
		t.Errorf("There were errors when querying for a domain: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
