package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/t-tashiro/yakiniku/account"
	"github.com/t-tashiro/yakiniku/db"
)

// Execute run migration
func Execute() {
	odb := db.ConnectDB()
	Migrate(odb)
	odb.Close()
}

// Migrate migretion
func Migrate(odb *gorm.DB) {
	odb.AutoMigrate(
		&account.Account{},
	)
}
