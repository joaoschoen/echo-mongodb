package user

import (
	// Project
	"API-ECHO-MONGODB/model"

	// Standard
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	// External
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// POST TESTS
var (
	badlyFormedJSON  = "{bad::form}"
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

// GET TESTS
var (
	goodResponse = "{\"Data\":{\"id\":\"someID\",\"email\":\"jon@doe.com\"}}\n"
)

func TestGetUser(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	METHOD := http.MethodPost
	URL := "/"
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var context echo.Context

	// TEST PREPARATION
	ValidID := func() {
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("someID")
	}
	NotFound := func() {
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("404")
	}

	// TESTS
	ValidID()
	if assert.NoError(t, GetUser(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, goodResponse, recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
	NotFound()
	if assert.NoError(t, GetUser(context)) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	}
}

func TestGetUserList(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	METHOD := http.MethodPost
	var URL string
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var context echo.Context

	// TEST PREPARATION
	FilterEmail := func() {
		q := make(url.Values)
		q.Set("email", "jon1@doe.com")
		URL = "/user/list?" + q.Encode()
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	CompletePage := func() {
		q := make(url.Values)
		q.Set("page", "1")
		URL = "/user/list?" + q.Encode()
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	IncompletePaging := func() {
		q := make(url.Values)
		q.Set("page", "2")
		URL = "/user/list?" + q.Encode()
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	NegativePaging := func() {
		q := make(url.Values)
		q.Set("page", "-2")
		URL = "/user/list?" + q.Encode()
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	// TESTS

	FilterEmail()
	if assert.NoError(t, GetUserList(context)) {
		expectedObjects := 1
		var objects model.GetUserListResponse
		err := json.Unmarshal(recorder.Body.Bytes(), &objects)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, expectedObjects, len(objects.Data))
	}
	// BadPaging
	CompletePage()
	if assert.NoError(t, GetUserList(context)) {
		expectedObjects := 2
		var objects model.GetUserListResponse
		json.Unmarshal(recorder.Body.Bytes(), &objects)
		assert.Equal(t, expectedObjects, len(objects.Data))
	}
	IncompletePaging()
	if assert.NoError(t, GetUserList(context)) {
		expectedObjects := 1
		var objects model.GetUserListResponse
		json.Unmarshal(recorder.Body.Bytes(), &objects)
		assert.Equal(t, expectedObjects, len(objects.Data))
	}
	NegativePaging()
	if assert.NoError(t, GetUserList(context)) {
		expectedPage := 0
		var response model.GetUserListResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.Equal(t, expectedPage, response.Paging.Page)
	}
}

// PUT TESTS

func TestPutUser(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	METHOD := http.MethodPost
	URL := "/user"
	var DATA *strings.Reader
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var context echo.Context

	// TEST PREPARATION
	NotFound := func() {
		user, err := json.Marshal(successfulUser)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("404")
	}
	EmptyObjectBodyTest := func() {
		user, err := json.Marshal(emptyUser)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("someID")
	}
	AlreadyInUseTest := func() {
		user, err := json.Marshal(alreadyInUseUser)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("someID")
	}
	SuccessTest := func() {
		user, err := json.Marshal(successfulUser)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("someID")
	}

	// TESTS
	NotFound()
	if assert.NoError(t, PutUser(context)) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	}

	EmptyObjectBodyTest()
	if assert.NoError(t, PutUser(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	AlreadyInUseTest()
	if assert.NoError(t, PutUser(context)) {
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	}

	SuccessTest()
	if assert.NoError(t, PutUser(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}

// DELETE TESTS

func TestDeleteUser(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	METHOD := http.MethodPost
	URL := "/user"
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var context echo.Context

	// TEST PREPARATION
	NotFound := func() {
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("404")
	}
	SuccessTest := func() {
		request = httptest.NewRequest(METHOD, URL, nil)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("someID")
	}

	// TESTS
	NotFound()
	if assert.NoError(t, DeleteUser(context)) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	}

	SuccessTest()
	if assert.NoError(t, DeleteUser(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}
