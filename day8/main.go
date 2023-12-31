package main

import (
	"fmt"
	"os"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Function to calculate LCM (Least Common Multiple)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate LCM of a list of integers
func LCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func ParseInput(input string) ([]byte, map[string][]string) {
	split := strings.Split(input, "\n\n")
	instructions := []byte(strings.Split(split[0], "\n")[0])
	nodes := make(map[string][]string, 0)

	for _, node_str := range strings.Split(split[1], "\n") {
		var node, left, right string
		split1 := strings.Split(node_str, " = ")

		node = split1[0]

		split2 := strings.Split(split1[1], ", ")
		left = split2[0][1:]
		right = split2[1][:len(split2[1])-1]

		nodes[node] = []string{left, right}
	}
	return instructions, nodes
}

func solve1(input string) int {
	instructions, nodes := ParseInput(input)

	pos := "AAA"
	i := 0
	for {
		ip := i % len(instructions)
		if instructions[ip] == 'L' {
			pos = nodes[pos][0]
		} else if instructions[ip] == 'R' {
			pos = nodes[pos][1]
		} else {
			panic("Unknown instruction")
		}
		if pos == "ZZZ" {
			break
		}
		i++
	}

	return i + 1
}

func solve2(input string) int {
	instructions, nodes := ParseInput(input)

	loops := make([]int, 0)
	for pos := range nodes {
		if []byte(pos)[2] != 'A' {
			continue
		}
		i := 0
		for {
			ip := i % len(instructions)
			if instructions[ip] == 'L' {
				pos = nodes[pos][0]
			} else if instructions[ip] == 'R' {
				pos = nodes[pos][1]
			} else {
				panic("Unknown instruction")
			}

			if []byte(pos)[2] == 'Z' {
				loops = append(loops, i+1)
				break
			}
			i++
		}
	}
	return LCM(loops)
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
