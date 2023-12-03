package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IsSymbol(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

func solve1(input string) int {
	lines := strings.Split(input, "\n")
	sol := 0

	for i, line := range lines {
		// since duplications may happen this needs to be a map
		nums := make(map[int]int)
		buf := make([]rune, 0)
		for _, c := range line {
			if IsSymbol(c) {
				if len(buf) > 0 {
					num, _ := strconv.Atoi(string(buf))
					nums[num] = nums[num] + 1
					buf = make([]rune, 0)
				}
			} else {
				buf = append(buf, c)
			}
		}
		if len(buf) > 0 {
			num, _ := strconv.Atoi(string(buf))
			nums[num] = nums[num] + 1
		}

		// merge prev line with line and next line
		merged := []rune(line)
		if i > 0 {
			prev := lines[i-1]
			for i, c := range prev {
				if IsSymbol(c) && c != '.' {
					merged[i] = c
				}
			}
		}
		if i < len(lines)-1 {
			next := lines[i+1]
			for i, c := range next {
				if IsSymbol(c) && c != '.' {
					merged[i] = c
				}
			}
		}

		candidates := strings.Split(string(merged), ".")
	loop:
		for _, candidate := range candidates {
			if candidate == "" {
				continue
			}
			for _, c := range candidate {
				if IsSymbol(c) {
					continue loop
				}
			}
			part_no, err := strconv.Atoi(candidate)
			if err == nil {
				nums[part_no] = nums[part_no] - 1
			}
		}
		for k, v := range nums {
			sol += v * k
		}
	}

	return sol
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")
	y := len(lines)
	x := len(lines[0])
	mat := make([]byte, y*x)
	sol := 0

	for i, line := range lines {
		for j, c := range line {
			mat[i*x+j] = byte(c)
		}
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			c := mat[i*x+j]
			if c == '*' {
				nums := make(map[int]bool, 0)
				adjacents := []int{i*x + j - 1, i*x + j + 1, (i-1)*x + j, (i+1)*x + j, (i-1)*x + j - 1, (i-1)*x + j + 1, (i+1)*x + j - 1, (i+1)*x + j + 1}
				for _, adj := range adjacents {
					if adj < 0 || adj >= x*y {
						continue
					}
					if mat[adj] >= '0' && mat[adj] <= '9' {
						go_left := 0
						for {
							if adj-go_left < 0 {
								break
							}
							if mat[adj-go_left] >= '0' && mat[adj-go_left] <= '9' {
								go_left++
							} else {
								break
							}
						}
						go_right := 0
						for {
							if adj+go_right > x*y {
								break
							}
							if mat[adj+go_right] >= '0' && mat[adj+go_right] <= '9' {
								go_right++
							} else {
								break
							}
						}
						n, _ := strconv.Atoi(string(mat[adj-go_left+1 : adj+go_right]))
						nums[n] = true
					}
				}
				// just one vertex
				if len(nums) < 2 {
					continue
				}
				gear := 1
				for n, _ := range nums {
					gear *= n
				}
				sol += gear
			}
		}
	}
	return sol
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
