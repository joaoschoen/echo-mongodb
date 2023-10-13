package model

// Auth Token response model
// @Description Access token generated from user and password
type AuthToken struct {
	Token string `json:"token"`
}
