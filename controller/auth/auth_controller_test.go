package auth

import (
	// Project
	"API-ECHO-MONGODB/model"

	// Standard
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// External
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// POST TESTS
var (
	emptyUser        = model.UnsafeUser{}
	alreadyInUseUser = model.UnsafeUser{
		Email:    "alreadyIn@use.com",
		Password: "BadExample",
	}
	successfulUser = model.UnsafeUser{
		Email:    "jon@doe.com",
		Password: "BadExample",
	}
)

func TestLogin(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	METHOD := http.MethodPost
	URL := "/user"
	var DATA *strings.Reader
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var context echo.Context

	// TEST PREPARATION
	EmptyObjectBodyTest := func() {
		user, err := json.Marshal(emptyUser)
		if err != nil {
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		return
	}
	AlreadyInUseTest := func() {
		user, err := json.Marshal(alreadyInUseUser)
		if err != nil {
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		return
	}
	SuccessTest := func() {
		user, err := json.Marshal(successfulUser)
		if err != nil {
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		return
	}

	// TESTS
	EmptyObjectBodyTest()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	AlreadyInUseTest()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	}

	SuccessTest()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}

// GET TESTS
var (
	goodResponse = "{\"Data\":{\"id\":\"someID\",\"email\":\"jon@doe.com\"}}\n"
)
