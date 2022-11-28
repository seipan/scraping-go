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

func CalcAvgRuntime(target func(), N int) time.Duration {
	var avgRuntime time.Duration
	for i := 0; i < N; i++ {
		fmt.Printf("\n%s time\n", strconv.Itoa(i+1))
		avgRuntime += measurer(target)
	}
	avgRuntime = avgRuntime / time.Duration(int64(N))
	return avgRuntime
}
