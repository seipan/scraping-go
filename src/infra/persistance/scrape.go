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

func CreateHistory(db *sql.DB, user_id string, article_id int) (*Article, error) {
	statement := "INSERT INTO histories (user_id, article_id) VALUES($1,$2) returning user_id, article_id, created_at"

	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	stmt.Exec()
	article := &Article{}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return article, nil
}
