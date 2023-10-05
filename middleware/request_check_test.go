package middleware

import (
	"API-ECHO-MONGODB/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	badlyFormedJSON = "{bad::form}"
	wellFormedJSON  = "{\"email\":\"jon@doe.com\",\"password\":\"BadExample\"}"
)

func TestCheckBody(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	TestServer.POST("/test", CheckBody(SampleFunc, model.UnsafeUser{}))
	METHOD := http.MethodPost
	URL := "/test"
	var DATA *strings.Reader
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	// TEST PREPARATION
	BadlyFormedJSON := func() {
		DATA = strings.NewReader(badlyFormedJSON)
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
	}
	WellFormedJSON := func() {
		DATA = strings.NewReader(wellFormedJSON)
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
	}
	// TESTS
	BadlyFormedJSON()
	TestServer.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	WellFormedJSON()
	TestServer.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
func SampleFunc(echo echo.Context) error {
	println("TEST::")
	var user model.UnsafeUser

	if err := echo.Bind(&user); err != nil {
		return echo.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	println(user.Email)
	println(user.Password)
	if user.Email == "jon@doe.com" && user.Password == "BadExample" {
		return echo.JSON(http.StatusOK, "Correctly parsing json")
	}
	return echo.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
}
