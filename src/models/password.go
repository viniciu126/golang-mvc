package models

// Passwd represents a struct to change password
type Passwd struct {
	New    string `json:"new"`
	Actual string `json:"actual"`
}
