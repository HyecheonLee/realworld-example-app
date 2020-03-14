package handler_test

import (
	"github.com/hyecheonlee/realworld-example-app/handler"
	"github.com/hyecheonlee/realworld-example-app/router"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	articlesRes = `{"articles":[{"slug":"article1-slug","title":"article1 title","description":"article1 description""body":"article1 body","tagList":["tag1","tag2"],"createdAt":"2016-02-18T03:22:56.637Z","updatedAt":"2016-02-18T03:48:35.824Z","favorited": false,"favoritesCount":0,"author":{"username":"user1","bio":"user1 bio","image":"http://realworld.io/user1.jpg","following":false}}],"articlesCount": 2}`
)

func TestListArticles(t *testing.T) {
	tearDown()
	setup()
	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/articles", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.New(db)

	if assert.NoError(t, h.Articles(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Regexp(t, articlesRes, rec.Body.String())
	}

}
