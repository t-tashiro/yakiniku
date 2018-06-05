package repository_test

import (
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/t-tashiro/yakiniku/account/repository"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func stubDB(dsn string) (*gorm.DB, sqlmock.Sqlmock) {
	var db *gorm.DB
	_, mock, err := sqlmock.NewWithDSN(dsn)
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err = gorm.Open("sqlmock", dsn)
	if err != nil {
		panic("Got an unexpected error.")
	}
	return db, mock
}

func TestList(t *testing.T) {
	db, mock := stubDB("list")
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(uuid.NewV4().String(), "mockaccount", "aaaa@example.com", "asdf0987")

	query := `SELECT * FROM "accounts" WHERE "accounts"."deleted_at" IS NULL`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	a := repository.NewAccountRepository(db)
	accounts, err := a.List()
	assert.NoError(t, err)
	assert.NotNil(t, accounts)
}
