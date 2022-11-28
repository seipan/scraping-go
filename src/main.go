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

	log.Println("1")

	articleChan := make(chan []presentation.Article, per_page*page)
	var wg sync.WaitGroup

	log.Println("2")

	for i := 0; i < page; i++ {
		wg.Add(1)
		go presentation.GetParallelQiitaArticle(i, per_page, articleChan)
		wg.Done()
	}

	log.Println("3")

	wg.Wait()

	log.Println("4")

	var ResArticle []*domain.Article

	for articles := range articleChan {
		for _, article := range articles {
			log.Println("GetArticle   " + article.Title)
			domArticle := utils.JsonToDomain(article)
			ResArticle = append(ResArticle, &domArticle)
		}
	}

	log.Println("4")

	log.Println("---------------------------")
}

func main() {
	per_page := 100
	//var avgRuntime time.Duration

	avgRuntime := utils.CalcAvgRuntime(GetParallelQiitaArticle, per_page)
	log.Println(avgRuntime)
}
