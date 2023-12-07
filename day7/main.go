package main

import (
	"day1/aoc/part1"
	"day1/aoc/part2"
	"fmt"
	"os"
)

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", part1.Solve(input), part2.Solve(input))
}
