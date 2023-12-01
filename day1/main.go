package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetNumber(s string) int {
	num := make([]int, 0)
	for _, char := range s {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		} else {
			num = append(num, n)
		}
	}
	return num[0]*10 + num[len(num)-1]
}
func solve1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += GetNumber(line)
	}
	return sum
}

func solve2(input string) int {
	nums := []*regexp.Regexp{
		regexp.MustCompile("one"),
		regexp.MustCompile("two"),
		regexp.MustCompile("three"),
		regexp.MustCompile("four"),
		regexp.MustCompile("five"),
		regexp.MustCompile("six"),
		regexp.MustCompile("seven"),
		regexp.MustCompile("eight"),
		regexp.MustCompile("nine"),
	}

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		ll := []rune(line)
		for index, r := range nums {
			matches := r.FindAllIndex([]byte(line), -1)
			for _, i := range matches {
				ll[i[0]] = rune('0' + index + 1)
			}
		}
		sum += GetNumber(string(ll))
	}
	return sum
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
