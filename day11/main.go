package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func ManhattanDistance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0]))) + int(math.Abs(float64(a[1]-b[1])))
}

func Empty(grid [][]rune) ([]int, []int) {
	r := make([]int, 0)
	c := make([]int, 0)

	cols := make([][]rune, len(grid[0]))
	for i := range cols {
		cols[i] = make([]rune, len(grid))
	}

	// expand vertically
	for j, line := range grid {
		if strings.ReplaceAll(string(line), ".", "") == "" {
			r = append(r, j)
		}
		for i, c := range line {
			cols[i][j] = c
		}
	}

	// expand horizontally
	for col, line := range cols {
		if strings.ReplaceAll(string(line), ".", "") == "" {
			c = append(c, col)
		}
	}

	return r, c

}
func ExpandGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, 0)
	cols := make([][]rune, len(grid[0]))
	for i := range cols {
		cols[i] = make([]rune, len(grid))
	}

	// expand vertically
	for j, line := range grid {
		if strings.ReplaceAll(string(line), ".", "") == "" {
			newGrid = append(newGrid, line)
		}
		newGrid = append(newGrid, line)
		for i, c := range line {
			cols[i][j] = c
		}
	}

	// expand horizontally
	d := 0
	for col, line := range cols {
		if strings.ReplaceAll(string(line), ".", "") == "" {
			for i := range newGrid {
				newGrid[i] = slices.Insert(newGrid[i], col+d, '.')
			}
			d++

		}
	}

	return newGrid

}

func solve1(input string) int {
	sum := 0
	stars := make(map[int][2]int, 0)
	id := 1
	grid := make([][]rune, 0)
	for y, line := range strings.Split(input, "\n") {
		grid = append(grid, make([]rune, len(line)))
		for x, char := range line {
			grid[y][x] = char
		}
	}
	grid = ExpandGrid(grid)

	for y, line := range grid {
		for x, char := range line {
			if char == '#' {
				stars[id] = [2]int{x, y}
				id++
			}
		}
	}

	for _, v1 := range stars {
		for _, v2 := range stars {
			sum += ManhattanDistance(v1, v2)
		}
	}

	return sum / 2
}

func solve2(input string) int {
	sum := 0
	stars := make(map[int][2]int, 0)
	id := 1
	grid := make([][]rune, 0)
	for y, line := range strings.Split(input, "\n") {
		grid = append(grid, make([]rune, len(line)))
		for x, char := range line {
			grid[y][x] = char
		}
	}
	for y, line := range grid {
		for x, char := range line {
			if char == '#' {
				stars[id] = [2]int{x, y}
				id++
			}
		}
	}

	space := 1000000 - 1
	rows, cols := Empty(grid)
	for k, v1 := range stars {
		delta := 0
		for _, c := range cols {
			if v1[0] > c {
				delta += space
			}

		}
		v1[0] += delta
		delta = 0

		for _, r := range rows {
			if v1[1] > r {
				delta += space
			}

		}
		v1[1] += delta
		stars[k] = v1
	}
	for _, v1 := range stars {
		for _, v2 := range stars {
			sum += ManhattanDistance(v1, v2)
		}
	}

	return sum / 2
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
