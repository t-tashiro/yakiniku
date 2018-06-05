package repository

import (
	"github.com/jinzhu/gorm"
	model "github.com/t-tashiro/yakiniku/account"
)

// accountRepository account repository
type accountRepository struct {
	Conn *gorm.DB
}

// NewAccountRepository mount account repository
func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		Conn: db,
	}
}

// AccountRepository account repository interfce
type AccountRepository interface {
	List() ([]*model.Account, error)
}

// List list accounts
func (m *accountRepository) List() ([]*model.Account, error) {
	var accounts = []*model.Account{}
	// SELECT * FROM "accounts" WHERE "accounts".deleted_as IS NULL
	err := m.Conn.Model(&accounts).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
