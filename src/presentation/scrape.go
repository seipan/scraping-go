package presentation

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/seipan/scraping-go/domain"
	"github.com/seipan/scraping-go/scrape"
)

func GetQiitaArticle(page int, per_page int) *domain.Article {
	endpoint := "http://qiita.com/api/v2/users/snaka/items"

	values := url.Values{}

	values.Set("page", strconv.Itoa(page))
	values.Add("per_page", strconv.Itoa(per_page))

	data, err := scrape.DoAPI("GET", endpoint, values, nil)
	if err != nil {
		log.Println("DoAPI Error() = %w", err)
		return nil
	}

	Respon := &domain.Article{}

	if err := json.Unmarshal(data, &Respon); err != nil {
		log.Println("DoAPI Error() = %w", err)
		return nil
	}

	return Respon

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
