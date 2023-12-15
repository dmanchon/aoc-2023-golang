package main

import (
	"fmt"
	"os"
	"strings"
)

func TransposeGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid[0]))
	for i := range newGrid {
		newGrid[i] = make([]string, len(grid))
	}

	for i, row := range grid {
		for j, val := range row {
			newGrid[j][i] = val
		}
	}
	return newGrid
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func BuildGrid(input string) [][]string {
	grid := make([][]string, 0)
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func solve1(input string) int {
	sum := 0

	g := BuildGrid(input)
	t := TransposeGrid(g)
	for _, row := range t {
		r := strings.Join(row, "")
		ss := strings.Split(r, "#")
		p := 0
		for _, s := range ss {
			for i := 0; i < strings.Count(s, "O"); i++ {
				sum += len(r) - i - p
			}
			p += len(s) + 1
		}
	}
	return sum
}

func solve2(input string) int {

	sum := 0

	g := BuildGrid(input)

	PrintGrid(g)
	Tilt(&g)
	PrintGrid(g)

	return sum
}

func Tilt(g *[][]string) {
	*g = TransposeGrid(*g)
	for i, row := range *g {
		r := strings.Join(row, "")
		ss := strings.Split(r, "#")
		p := 0
		for _, s := range ss {
			for k := 0; k < strings.Count(s, "O"); k++ {
				j := len(r) - k - p - 1
				(*g)[i][j] = "O"
			}
			p += len(s) + 1
		}
	}
	*g = TransposeGrid(*g)
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
