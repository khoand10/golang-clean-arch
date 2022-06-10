package database

import (
	"fmt"
	"golang-clean-arch/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Sqlite struct {
	DB *gorm.DB
}

var db *gorm.DB
var err error

func InitSqlite() (*gorm.DB, func(), error) {
	db, err = gorm.Open(sqlite.Open("db-storage/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}

	err := migrateTable(db)
	if err != nil {
		log.Fatalf("model.migrateTable err: %v", err)
	}

	sqlDB, err := db.DB()
	cleanFunc := func() {
		err := sqlDB.Close()
		if err != nil {
			fmt.Errorf("Gorm db close error: %s", err.Error())
		}
	}
	return db, cleanFunc, err
}

func migrateTable(db *gorm.DB) error {
	//var allModels = []interface{}{new(model.User)}
	return db.AutoMigrate(model.User{})
}
