package main

import (
	"fmt"
	"math"
)

func RootsDiff(t, d int) int {
	a, b, c := 1.0, float64(t), float64(d)
	x1 := (b + math.Sqrt(b*b-4*c)) / (2 * a)
	x2 := (b - math.Sqrt(b*b-4*c)) / (2 * a)
	return int(x1) - int(x2)
}

func solve1(input [][]int) int {
	sum := 1
	for _, v := range input {
		t, d := v[0], v[1]
		sum *= RootsDiff(t, d)
	}
	return sum
}

func main() {
	input := [][]int{{41, 214}, {96, 1789}, {88, 1127}, {94, 1055}}
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), RootsDiff(41968894, 214178911271055))
}
