package persistance

import (
	"database/sql"
	"log"

	"github.com/seipan/scraping-go/domain"
)

func CreateArticle(db *sql.DB, article *domain.Article) (*domain.Article, error) {
	statement := "INSERT INTO articles (id, title, body, url) VALUES($1,$2,$3,$4)"

	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	stmt.Exec()
	resarticle := &domain.Article{}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return resarticle, nil
}
