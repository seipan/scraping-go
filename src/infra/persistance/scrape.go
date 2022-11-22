package persistance

import (
	"database/sql"
	"log"
)

type Article struct {
	Id     int
	Title  string
	Likes  int
	Url    string
	Author string
	Kind   string
}

func CreateArticle(db *sql.DB, article *Article) (*Article, error) {
	statement := "INSERT INTO articles (id, title, body, url) VALUES($1,$2,$3,$4)"

	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	stmt.Exec()
	resarticle := &Article{}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return resarticle, nil
}
