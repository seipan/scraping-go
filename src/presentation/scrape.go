package presentation

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/seipan/scraping-go/domain"
	"github.com/seipan/scraping-go/scrape"
)

func GetQiitaArticle(page int, per_page int, articleChan chan *domain.Article) {
	endpoint := "http://qiita.com/api/v2/users/snaka/items"

	values := url.Values{}

	values.Set("page", strconv.Itoa(page))
	values.Add("per_page", strconv.Itoa(per_page))

	data, err := scrape.DoAPI("GET", endpoint, values, nil)
	if err != nil {
		panic(err)
	}

	Respon := &domain.Article{}

	if err := json.Unmarshal(data, &Respon); err != nil {
		panic(err)
	}

	articleChan <- Respon
}
func GetParallelQiitaArticle(page int, per_page int, articleChan chan *domain.Article) {
	endpoint := "http://qiita.com/api/v2/users/snaka/items"

	values := url.Values{}
	values.Set("page", strconv.Itoa(page))
	values.Add("per_page", strconv.Itoa(per_page))

	ch := make(chan []byte)

	go scrape.DoAPIParallel("GET", endpoint, values, nil, ch)
	data := <-ch

	Respon := &domain.Article{}

	if err := json.Unmarshal(data, &Respon); err != nil {
		panic(err)
	}

	articleChan <- Respon
}

type Article struct {
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
	ID         string    `json:"id"`
	LikesCount int       `json:"likes_count"`
	Title      string    `json:"title"`
	URL        string    `json:"url"`
	User       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}
