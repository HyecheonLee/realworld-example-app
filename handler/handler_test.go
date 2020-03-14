package handler_test

import (
	"github.com/hyecheonlee/realworld-example-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func testDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func AutoMigrate() {
	if err := db.AutoMigrate(&models.User{}).Error; err != nil {
		log.Fatal(err)
	}
}

func dropTestDB() {
	_ = db.Close()
	if err := os.Remove("./../realworld_test.db"); err != nil {
		log.Fatal(err)
	}
}
func authHeader(token string) string {
	return "Token " + token
}
func setup() {
	db = testDB()
	AutoMigrate()
	loadFixtures()
}
func tearDown() {
	dropTestDB()
}
func loadFixtures() error {
	bio := "user1 bio"
	image := "http://realworld.io/user1.jpg"
	u := models.User{
		Username: "user1",
		Email:    "user1@realworld.io",
		Bio:      &bio,
		Image:    &image,
	}

	if err := u.HashPassword("secret"); err != nil {
		return err
	}
	if err := db.Create(&u).Error; err != nil {
		return nil
	}
	return nil
}
