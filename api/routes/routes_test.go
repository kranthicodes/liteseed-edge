package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liteseed/bungo/internal/database"
	"github.com/liteseed/bungo/internal/store"
)

// NewApiTest returns new API test helper.
func NewApiTest() (*gin.Engine, *Routes) {
	gin.SetMode(gin.TestMode)

	db := database.NewSqliteDatabase("./tmp/sqlite")
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	store := store.New("pebble")

	a := New(db, store)

	r := gin.Default()

	r.GET("/status", a.GetStatus)
	r.POST("/data", a.PostData)

	return r, a
}

// PerformRequest runs an API request with an empty request body.
// See https://medium.com/@craigchilds94/testing-gin-json-responses-1f258ce3b0b1
func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

// PerformRequestWithBody runs an API request with the request body as a string.
func PerformRequestWithBody(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)
	req, _ := http.NewRequest(method, path, reader)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
