package main

import (
	"log"
	"sync"

	"github.com/seipan/scraping-go/infra/db"
	"github.com/seipan/scraping-go/infra/persistance"
	"github.com/seipan/scraping-go/presentation"
	"github.com/seipan/scraping-go/utils"
)

func GetParallelQiitaArticle() []presentation.Article {
	log.Println("---------------------------")
	page := 100
	per_page := 100

	articleChan := make(chan []presentation.Article, per_page*page)
	var wg sync.WaitGroup

	for i := 0; i < page; i++ {
		wg.Add(1)
		go presentation.GetParallelQiitaArticle(i, per_page, articleChan)
		wg.Done()
	}

	wg.Wait()

	var ResArticle []presentation.Article

	for articles := range articleChan {
		for _, article := range articles {
			log.Println("GetArticle   " + article.Title)
			ResArticle = append(ResArticle, article)
		}
	}

	log.Println("---------------------------")

	return ResArticle
}

func InsertQiitaArticle(articles []presentation.Article) {
	db, err := db.NewDriver()
	if err != nil {
		panic(err)
	}

	for _, article := range articles {
		domarticle := utils.JsonToDomain(article)
		persistance.CreateArticle(db, domarticle)
	}
}

func main() {
	time, article := utils.Measurer(GetParallelQiitaArticle)
	log.Println(time)

	InsertQiitaArticle(article)
}
