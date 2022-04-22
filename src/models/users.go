package models

import (
	"api/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// User represents a user
type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Passwd    string `json:"passwd,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

// Prepare validate and format user received
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if step == "signup" {
		if user.Name == "" {
			return errors.New("name is required")
		}

		if user.Email == "" {
			return errors.New("email is required")
		}

		if user.Nick == "" {
			return errors.New("nick is required")
		}

		if user.Passwd == "" {
			return errors.New("password is required")
		}
	}

	if user.Email != "" {
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Email format invalid")
		}
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "signup" {
		passwordHashed, err := security.Hash(user.Passwd)
		if err != nil {
			return err
		}

		user.Passwd = string(passwordHashed)
	}

	return nil
}
