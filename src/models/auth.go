package models

// AuthData contains jwt token and user ID
type AuthData struct {
	ID    string `json="id"`
	Token string `json="token"`
}
