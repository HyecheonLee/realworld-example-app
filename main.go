package main

import (
	"github.com/hyecheonlee/realworld-example-app/database"
	"github.com/hyecheonlee/realworld-example-app/handler"
	"github.com/hyecheonlee/realworld-example-app/router"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	db := database.New()
	db.AutoMigrate()
	h := handler.New(db)
	r := router.New()

	r.POST("/api/users/login", h.Login)
	r.POST("/api/users", h.Register)
	r.GET("/api/user", h.CurrentUser)
	r.PUT("/api/user", h.UpdateUser)

	r.GET("/api/profiles/:username", h.GetProfile)
	r.POST("/api/profiles/:username/follow", h.Follow)
	r.DELETE("/api/profiles/:username/follow", h.Unfollow)

	r.GET("/api/articles", h.Articles)
	r.GET("/api/articles/feed", h.Feed)
	r.GET("/api/articles/:slug", h.GetArticles)
	r.POST("/api/articles", h.CreateArticle)
	r.PUT("/api/articles/:slug", h.UpdateArticle)
	r.DELETE("/api/articles/:slug", h.DeleteArticle)

	r.POST("/api/articles/:slug/comments", h.AddComment)
	r.GET("/api/articles/:slug/comments", h.GetComments)
	r.DELETE("/api/articles/:slug/comments/:id", h.DeleteComment)

	r.POST("/api/articles/:slug/favorite", h.Favorite)
	r.DELETE("/api/articles/:slug/favorite", h.Unfavorite)

	r.GET("/api/tags", h.Tags)
}
