package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Sets up mockDb
func SetUp(t *testing.T) (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New()

	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDb,
	}))

	return db, mockDb, mock
}
