package main

import (
	"github.com/hyecheonlee/realworld-example-app/db"
	"github.com/hyecheonlee/realworld-example-app/router"
	"github.com/hyecheonlee/realworld-example-app/store"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)
	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)

	r.Logger.Fatal(r.Start("127.0.0.1:1323"))
}
