package database

import (
	"go-backoffice-seller-api/src/config"

	"github.com/jinzhu/gorm"
)

type IDatabaseEngine interface {
	// TODO read from config file
	GetDatabase(config config.Database, env string) *gorm.DB
	RunMigration()
}
