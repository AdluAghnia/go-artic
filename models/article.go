package models

import (
	"database/sql"

	"github.com/AdluAghnia/go-artic/db"
	"github.com/google/uuid"
)

type Article struct {
	ID      string
	Title   string
	Content string
	Image   string
}

func NewArticle(title, content, image string) *Article {
	id := uuid.New().String()
	return &Article{
		ID:      id,
		Title:   title,
		Content: content,
		Image:   image,
	}
}

func (a *Article) SaveArticle(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO article (id, title, content, image) VALUES (? ,?, ?, ?)", a.ID, a.Title, a.Content, a.Image)
	if err != nil {
		return err
	}

	return nil
}

func GetArticles() ([]*Article, error) {
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, title, content, image FROM article")
	if err != nil {
		return nil, err
	}

	var articles []*Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Image); err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}
