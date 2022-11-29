package utils

import (
	"fmt"
	"strconv"
	"time"
)

func measurer(fnc func()) time.Duration {
	fmt.Println("start scraping")
	start := time.Now()
	fnc()
	end := time.Now()
	return end.Sub(start)
}

func CalcTime(kansu func(), N int) time.Duration {
	var avgtime time.Duration

	for i := 0; i < N; i++ {
		fmt.Printf("\n%s time\n", strconv.Itoa(i+1))
		avgtime += measurer(kansu)
	}
	avgtime = avgtime / time.Duration(int64(N))
	return avgtime
}
