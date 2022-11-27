package persistance

import (
	"database/sql"
	"log"

	"github.com/seipan/scraping-go/utils"
)

func CreateArticle(db *sql.DB, article *utils.Article) (*utils.Article, error) {
	statement := "INSERT INTO articles (id, title, body, url) VALUES($1,$2,$3,$4)"

	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	stmt.Exec()
	resarticle := &utils.Article{}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return resarticle, nil
}
