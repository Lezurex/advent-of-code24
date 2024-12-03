package solutions

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func SolveDay01(input []string) {
	a := make([]int, 0)
	b := make([]int, 0)
	for _, l := range input {
		parts := strings.Split(l, "   ")
		x, err := strconv.Atoi(parts[0])
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		a = append(a, x)
		b = append(b, y)
	}

	slices.Sort(a)
	slices.Sort(b)
	distances := make([]int, 0)
	score := 0

	for i, x := range a {
		y := b[i]
		distances = append(distances, int(math.Abs(float64(x-y))))
		ocurrences := 0
		for _, y2 := range b {
			if x == y2 {
				ocurrences += 1
			}
		}
		score += x * ocurrences
	}

	sum := 0
	for _, d := range distances {
		sum += d
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Score: %d\n", score)
}
