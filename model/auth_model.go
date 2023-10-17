package model

// Auth Token response model
// @Description Stringified JWT access token
type AuthToken struct {
	Token string `json:"token"`
}
