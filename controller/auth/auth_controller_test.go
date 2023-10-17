package auth

import (
	// Project
	"API-ECHO-MONGODB/model"
	"API-ECHO-MONGODB/mongodb"
	"API-ECHO-MONGODB/test"
	"context"

	// Standard
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// External
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DUMMY DATA
var (
	emptyJsonString = "{}"
	emptyUser       = model.UnsafeUser{
		Email:    "",
		Password: "",
	}
	notFound = model.UnsafeUser{
		Email:    "notFound@test.com",
		Password: "notFound",
	}
	wrongPassword = model.UnsafeUser{
		Email:    "login@test.com",
		Password: "wrongPassword",
	}
	successfulUser = model.UnsafeUser{
		Email:    "login@test.com",
		Password: "loginTest",
	}
	alreadyInUseUser = model.UnsafeUser{
		Email:    "alreadyIn@use.com",
		Password: "BadExample",
	}
)

func TestRegister(t *testing.T) {
	// SETUP
	TestServer := echo.New()
	METHOD := http.MethodPost
	URL := "/user"
	var DATA *strings.Reader
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var context echo.Context

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
	}

	// TESTS
	EmptyObjectBodyTest()
	if assert.NoError(t, Register(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	AlreadyInUseTest()
	if assert.NoError(t, Register(context)) {
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	}

	SuccessTest()
	if assert.NoError(t, Register(context)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}

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
	EmptyJsonString := func() {
		user, err := json.Marshal(emptyJsonString)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	EmptyUser := func() {
		user, err := json.Marshal(emptyUser)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	NotFound := func() {
		user, err := json.Marshal(notFound)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	WrongPassword := func() {
		user, err := json.Marshal(wrongPassword)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}
	SuccessfulLogin := func() {
		user, err := json.Marshal(successfulUser)
		if err != nil {
			panic(err)
		}
		DATA = strings.NewReader(string(user))
		request = httptest.NewRequest(METHOD, URL, DATA)
		request.Header.Set("Content-Type", "application/json")
		recorder = httptest.NewRecorder()
		context = TestServer.NewContext(request, recorder)
	}

	// TESTS
	EmptyJsonString()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
	NotFound()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
	WrongPassword()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
	EmptyUser()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
	SuccessfulLogin()
	if assert.NoError(t, Login(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}

// TEST SETTINGS

var client *mongo.Client

func setup() {
	test.LoadTestEnv()
	client = mongodb.Connect()
}

func teardown() {
	// Delete user created during tests
	mongodb.DeleteOne("user", bson.D{{Key: "email", Value: successfulUser.Email}})
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}
