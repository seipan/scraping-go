package utils

import (
	"github.com/seipan/scraping-go/domain"
	"github.com/seipan/scraping-go/presentation"
)

func JsonToDomain(article presentation.Article) domain.Article {
	resUser := domain.User{}
	resUser.ID = article.User.ID
	resUser.Name = article.User.Name

	return domain.Article{
		Body:      article.Body,
		Title:     article.Title,
		CreatedAt: article.CreatedAt,
		URL:       article.URL,
		User:      resUser,
	}
}
