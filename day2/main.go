package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(input string) int {
	sol := 0
firstloop:
	for _, line := range strings.Split(input, "\n") {
		split1 := strings.Split(line, ":")
		split2 := strings.Split(split1[0], " ")
		game_id, _ := strconv.Atoi(string(split2[1]))
		for _, grab := range strings.Split(split1[1], ";") {
			result := make(map[string]int)

			split3 := strings.Split(grab, ",")
			for _, cube := range split3 {
				split4 := strings.Split(cube, " ")
				n, _ := strconv.Atoi(split4[1])
				color := split4[2]
				result[color] += n
			}
			if result["red"] > 12 || result["green"] > 13 || result["blue"] > 14 {
				continue firstloop
			}
		}
		sol += game_id

	}
	return sol
}

func solve2(input string) int {
	sol := 0
	for _, line := range strings.Split(input, "\n") {
		split1 := strings.Split(line, ":")
		result := make(map[string]int)
		for _, grab := range strings.Split(split1[1], ";") {

			split3 := strings.Split(grab, ",")
			for _, cube := range split3 {
				split4 := strings.Split(cube, " ")
				n, _ := strconv.Atoi(split4[1])
				color := split4[2]
				if result[color] < n {
					result[color] = n
				}
			}
		}
		sol += result["red"] * result["green"] * result["blue"]
	}
	return sol
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
