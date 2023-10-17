package auth

import (
	"API-ECHO-MONGODB/model"
	"API-ECHO-MONGODB/mongodb"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// @Summary		Login
// @Description	Receives user email and password, returns token
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param 			request	body		model.UnsafeUser	true	"Login information"
// @Success		200		{object}	model.AuthToken
// @Failure		400 "Bad Request"
// @Failure		500	"Internal server error"
// @Router			/auth/login [post]
func Login(echo echo.Context) error {
	// BODY
	var loginInfo model.UnsafeUser
	if err := echo.Bind(&loginInfo); err != nil {
		return echo.JSON(http.StatusBadRequest, "Bad username and password combination")
	}

	// Empty data
	if loginInfo.Email == "" || loginInfo.Password == "" {
		return echo.JSON(http.StatusBadRequest, "Bad username and password combination")
	}

	// GET USER
	var user model.UnsafeUser
	if err := mongodb.FindOne("user", bson.D{{Key: "email", Value: loginInfo.Email}}).Decode(&user); err != nil {
		return echo.JSON(http.StatusBadRequest, "Bad username and password combination")
	}

	// BCRIPT
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		return echo.JSON(http.StatusBadRequest, "Bad username and password combination")
	}

	// INSERT
	TOKEN_SECRET := os.Getenv("TOKEN_SECRET")
	var token *jwt.Token
	var signedToken string
	token = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "echo-api",
			"sub": user.Email,
			"aud": "user",
			"exp": time.Now().AddDate(0, 0, 1).String(),
		})
	if signed, err := token.SignedString([]byte(TOKEN_SECRET)); err != nil {
		return echo.JSON(http.StatusInternalServerError, nil)
	} else {
		signedToken = signed
	}

	// BUILD RESPONSE
	response := model.AuthToken{
		Token: signedToken,
	}
	return echo.JSON(http.StatusOK, response)
}

// @Summary		Create new user
// @Description	Receives user email and password, returns UUID
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			email	path		string	true	"User email"
// @Success		201		{object}	model.PostUserResponse
// @Failure		400 "Bad Request"
// @Failure		422 "Email already in use"
// @Failure		500	"Internal server error"
// @Router			/auth/register [post]
func Register(echo echo.Context) error {
	// BODY
	var user model.UnsafeUser
	if err := echo.Bind(&user); err != nil {
		return echo.JSON(http.StatusBadRequest, "Error while parsing received data")
	}

	// Empty data
	if user.Email == "" || user.Password == "" {
		return echo.JSON(http.StatusBadRequest, "Error while parsing received data")
	}

	// CHECK DUPLICATE
	var duplicate model.UnsafeUser
	if err := mongodb.FindOne("user", bson.D{{Key: "email", Value: user.Email}}).Decode(&duplicate); err == nil {
		return echo.JSON(http.StatusUnprocessableEntity, "Email already in use")
	}

	// BCRIPT
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		panic(err)
	}

	// INSERT
	insert, err := mongodb.InsertOne("user", bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: hash},
	})
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, nil)
	}

	// BUILD RESPONSE
	response := model.PostUserResponse{
		ID: insert.InsertedID,
	}
	return echo.JSON(http.StatusCreated, response)
}
