package main

import (
	"log"
	"sync"

	"github.com/seipan/scraping-go/domain"
	"github.com/seipan/scraping-go/presentation"
)

func GetParallelQiitaArticle(page int, per_page int) *[]domain.Article {
	log.Println("---------------------------")
	page = 100

	articleChan := make(chan []presentation.Article, per_page*page)
	var wg sync.WaitGroup

	for i := 0; i < page; i++ {
		wg.Add(1)
		go presentation.GetParallelQiitaArticle(i, per_page, articleChan)
		wg.Done()
	}

	wg.Wait()

	var ResArticle []*domain.Article

	for articles := range articleChan {
		for _, article := range articles {
			ResArticle = append(ResArticle, articles)

		}
	}
}

func main() {
	log.Println("Hello")
}
