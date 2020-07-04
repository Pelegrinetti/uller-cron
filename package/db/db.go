package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

// New create a database instance
func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./uller.db")

	if err != nil {
		logrus.WithError(err).Panic("Error while connecting with database.")
	}

	return db
}

func Migrate(db *gorm.DB) {
    logrus.Info("Migrating tables...")
	db.AutoMigrate(&Metric{})
    logrus.Info("All tables migrated!")
}
