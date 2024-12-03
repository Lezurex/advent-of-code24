package solutions

import (
	"fmt"
	"regexp"
	"strconv"
)

func SolveDay03(input []string) {
	result := 0
	do := true
	for _, l := range input {
		re := regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
		for _, m := range re.FindAllStringSubmatch(l, -1) {
			if m[0] == "do()" {
				do = true
				continue
			}
			if m[0] == "don't()" {
				do = false
				continue
			}
			if !do {
				continue
			}
			a, err := strconv.Atoi(m[1])
			b, err := strconv.Atoi(m[2])
			if err != nil {
				panic(err)
			}
			result += a * b
		}
	}

	fmt.Printf("Result: %d\n", result)
}
