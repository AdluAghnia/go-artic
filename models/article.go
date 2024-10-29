package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Article struct {
	id      string
	title   string
	content string
	image   string
}

func NewArticle(title, content, image string) *Article {
	id := uuid.New().String()
	return &Article{
		id:      id,
		title:   title,
		content: content,
		image:   image,
	}
}

func (a *Article) SaveArticle(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO article (id, title, content, image) VALUES (? ,?, ?, ?)", a.id, a.title, a.content, a.image)
	if err != nil {
		return err	
	}

	return nil
}

