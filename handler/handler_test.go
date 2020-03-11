package handler_test

import (
	"fmt"
	"github.com/hyecheonlee/realworld-example-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	db = testDB()
	AutoMigrate()
	exitVal := m.Run()
	dropTestDB(db)
	os.Exit(exitVal)
}

func testDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func AutoMigrate() {
	db.AutoMigrate(&models.User{})
}
func dropTestDB(db *gorm.DB) error {
	db.Close()
	err := os.Remove("./../realworld_test.db")
	return err
}
