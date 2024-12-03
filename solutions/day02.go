package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveDay02(reports []string) {
	count := 0
	for _, report := range reports {
		levels := make([]int, 0)
		for _, level := range strings.Split(report, " ") {
			l, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			levels = append(levels, l)
		}

		problems := 0
		ascending := levels[0] < levels[1]
		for i, l := range levels[1:] {
			if levels[i] < l != ascending {
				problems++
				continue
			}
			if ascending && l-levels[i] > 3 {
				problems++
				continue
			}
			if !ascending && levels[i]-l > 3 {
				problems++
				continue
			}
			if levels[i]-l == 0 {
				problems++
				continue
			}
		}
		if problems <= 1 {
			count++
		}
	}

	fmt.Printf("Safe: %d\n", count)
}
