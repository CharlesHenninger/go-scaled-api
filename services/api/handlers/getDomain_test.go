package handlers

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetDomain_catchall(t *testing.T) {
	// mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	row := sqlmock.NewRows([]string{"id", "name", "events", "bounced"}).
		AddRow(1, "example1.com", 1500, false)
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT \* FROM domains WHERE name\=\$1`).WithArgs("example1.com").WillReturnRows(row)
	mockDB.Begin()

	// assertions
	results, err := getDomain(sqlxDB, "example1.com")
	if err != nil {
		t.Errorf("There were errors when querying for a domain: %s", err)
	}
	assert.Equal(t, "catch-all", results.IsCatchall)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetDomain_Not_catchall(t *testing.T) {
	// mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	row := sqlmock.NewRows([]string{"id", "name", "events", "bounced"}).
		AddRow(2, "example2.com", 10, true)
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT \* FROM domains WHERE name\=\$1`).WithArgs("example2.com").WillReturnRows(row)
	mockDB.Begin()

	// assertions
	results, err := getDomain(sqlxDB, "example2.com")
	if err != nil {
		t.Errorf("There were errors when querying for a domain: %s", err)
	}
	assert.Equal(t, "not catch-all", results.IsCatchall)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetDomain_unknown(t *testing.T) {
	// mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	row := sqlmock.NewRows([]string{"id", "name", "events", "bounced"}).
		AddRow(3, "example3.com", 10, false)
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT \* FROM domains WHERE name\=\$1`).WithArgs("example3.com").WillReturnRows(row)
	mockDB.Begin()

	// assertions
	results, err := getDomain(sqlxDB, "example3.com")
	if err != nil {
		t.Errorf("There were errors when querying for a domain: %s", err)
	}
	assert.Equal(t, "unknown", results.IsCatchall)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
