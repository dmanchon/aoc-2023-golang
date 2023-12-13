package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckConstraint(s string, n int, lens []int) bool {
	g := strings.Split(s, ".")

	// remove empty strings
	for i := 0; i < len(g); i++ {
		if g[i] == "" {
			g = append(g[:i], g[i+1:]...)
			i--
		}
	}

	if len(g) != n {
		return false
	}

	for i, v := range g {
		if v == "" {
			continue
		}
		if len(v) != lens[i] {
			return false
		}
	}
	return true
}

var cache = make(map[string][]string)

// given a string, return all possible combinations of the string
// substituting the characters '?' with combinations of '#' and '.'
func GetCombinations(s string, subs []byte) []string {
	if len(s) == 0 {
		return []string{""}
	}

	if v, ok := cache[s]; ok {
		return v
	}

	var res []string

	for _, v := range GetCombinations(s[1:], subs) {
		if s[0] == '?' {
			for _, sub := range subs {
				res = append(res, string(sub)+v)
			}
		} else {
			res = append(res, string(s[0])+v)
		}
	}
	cache[s] = res
	return res
}

func solve1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)

		s1 := strings.Split(line, " ")
		s2 := strings.Split(s1[1], ",")
		lens := make([]int, len(s2))
		for i, v := range s2 {
			lens[i], _ = strconv.Atoi(v)
		}

		for _, patt := range GetCombinations(s1[0], []byte{'#', '.'}) {
			if CheckConstraint(patt, len(lens), lens) {
				sum++
			}
		}
	}
	return sum
}

func solve2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
		s1 := strings.Split(line, " ")
		s2 := strings.Split(s1[1], ",")
		lens := make([]int, len(s2))
		for i, v := range s2 {
			lens[i], _ = strconv.Atoi(v)
		}

		for _, patt := range GetCombinations(s1[0], []byte{'#', '.'}) {

			if CheckConstraint(patt, len(lens), lens) {
				sum++
			}
		}

	}
	return sum

}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve1(input))
}
