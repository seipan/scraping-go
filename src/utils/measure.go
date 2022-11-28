package utils

import (
	"fmt"
	"time"
)

func measurer(fnc func()) time.Duration {
	fmt.Println("start scraping")
	start := time.Now()
	fnc()
	end := time.Now()
	return end.Sub(start)
}
