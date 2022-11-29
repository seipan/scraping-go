package utils

import (
	"fmt"
	"time"

	"github.com/seipan/scraping-go/presentation"
)

func Measurer(fnc func() []presentation.Article) (time.Duration, []presentation.Article) {
	fmt.Println("start scraping")
	start := time.Now()
	article := fnc()
	end := time.Now()
	return end.Sub(start), article
}
