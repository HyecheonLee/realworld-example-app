package main

import (
	"github.com/hyecheonlee/realworld-example-app/database"
	"github.com/hyecheonlee/realworld-example-app/handler"
	"github.com/hyecheonlee/realworld-example-app/router"
)

func main() {
	db := database.New()
	database.AutoMigrate(db)
	h := handler.New(db)
	r := router.New()
	h.RegisterRoutes(r)
	r.Logger.Fatal(r.Start("127.0.0.1:1323"))
}
