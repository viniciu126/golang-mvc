package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a user(Autor) Post
type Post struct {
	ID        uint64    `json:"id,omitempty`
	Title     string    `json:"title,omitempty`
	Content   string    `json:"content,omitempty`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
}

// Prepare validate and format post
func (post *Post) Prepare() error {
	if err := post.validar(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validar() error {
	if post.Title == "" {
		return errors.New("Title is required")
	}

	if post.Content == "" {
		return errors.New("Content is required")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
