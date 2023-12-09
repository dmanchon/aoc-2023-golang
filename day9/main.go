package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diff_series(series []int) []int {
	diff := make([]int, 0)
	for i := 0; i < len(series)-1; i++ {
		diff = append(diff, series[i+1]-series[i])
	}
	return diff
}

func all_zeroes(series []int) bool {
	for _, n := range series {
		if n != 0 {
			return false
		}
	}
	return true
}

func find_next(line string) int {
	series := make([][]int, 0)
	ts := make([]int, 0)

	for _, c := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(c)
		ts = append(ts, n)
	}
	series = append(series, ts)

	for {
		series = append(series, diff_series(series[len(series)-1]))
		if all_zeroes(series[len(series)-1]) {
			break
		}
	}

	inc := 0
	for i := len(series) - 1; i >= 0; i-- {
		last := series[i][len(series[i])-1]
		inc = inc + last
	}

	return inc
}

func find_prev(line string) int {
	series := make([][]int, 0)
	ts := make([]int, 0)

	for _, c := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(c)
		ts = append(ts, n)
	}
	series = append(series, ts)

	for {
		series = append(series, diff_series(series[len(series)-1]))
		if all_zeroes(series[len(series)-1]) {
			break
		}
	}

	inc := 0
	for i := len(series) - 1; i >= 0; i-- {
		last := series[i][0]
		inc = last - inc
	}

	return inc
}

func solve1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += find_next(line)
	}
	return sum
}

func solve2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += find_prev(line)
	}
	return sum
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
