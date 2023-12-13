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

var cache = make(map[int][]string)

// Get all combinations of elements
func GetCombinations(l int, elements []byte) []string {
	if v, ok := cache[l]; ok {
		return v
	}
	if l == 1 {
		res := make([]string, len(elements))
		for i, v := range elements {
			res[i] = string(v)
		}
		return res
	}

	res := make([]string, 0)
	for _, v := range elements {
		for _, s := range GetCombinations(l-1, elements) {
			res = append(res, string(v)+s)
		}
	}
	cache[l] = res
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

		for _, s := range GetCombinations(strings.Count(s1[0], "?"), []byte{'#', '.'}) {
			patt := s1[0]
			for _, v := range s {
				patt = strings.Replace(patt, "?", string(v), 1)
			}
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

		for _, s := range GetCombinations(strings.Count(s1[0], "?"), []byte{'#', '.'}) {
			patt := s1[0]
			for _, v := range s {
				patt = strings.Replace(patt, "?", string(v), 1)
			}
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
