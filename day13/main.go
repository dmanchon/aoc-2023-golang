package main

import (
	"fmt"
	"os"
	"strings"
)

func FindMirrorLine(line string) (int, error) {
	axis := -1
	rows := strings.Split(line, "\n")

	for i := 1; i < len(rows); i++ {
		for j := 1; j < i+1; j++ {
			if i+j-1 >= len(rows) {
				break
			}
			if rows[i-j] == rows[i+j-1] {
				axis = i
			} else {
				axis = -1
				break
			}
		}
		if axis != -1 {
			return axis, nil
		}
	}

	return axis, fmt.Errorf("no mirror line found")
}

func TransposeGrid(grid string) string {
	rows := strings.Split(grid, "\n")
	var cols []string
	for i := 0; i < len(rows[0]); i++ {
		var col string
		for j := 0; j < len(rows); j++ {
			col += string(rows[j][i])
		}
		cols = append(cols, col)
	}
	return strings.Join(cols, "\n")
}

func solve1(input string) int {
	sum := 0
	for _, grid := range strings.Split(input, "\n\n") {
		axis, ok := FindMirrorLine(grid)
		if ok == nil {
			sum += axis * 100
		}
		t := TransposeGrid(grid)
		axis, ok = FindMirrorLine(t)
		if ok == nil {
			sum += axis
		}
	}

	return sum
}

func solve2(input string) int {

	sum := 0

	return sum
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)

	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
