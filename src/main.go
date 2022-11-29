package main

import (
	"log"
	"sync"

	"github.com/seipan/scraping-go/domain"
	"github.com/seipan/scraping-go/presentation"
	"github.com/seipan/scraping-go/utils"
)

func GetParallelQiitaArticle() {
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

	var ResArticle []*domain.Article

	for articles := range articleChan {
		for _, article := range articles {
			log.Println("GetArticle   " + article.Title)
			domArticle := utils.JsonToDomain(article)
			ResArticle = append(ResArticle, &domArticle)
		}
	}

	log.Println("---------------------------")
}

func main() {
	per_page := 100
	//var avgRuntime time.Duration

	avgRuntime := utils.CalcTime(GetParallelQiitaArticle, per_page)
	log.Println(avgRuntime)
}
