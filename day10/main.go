package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Dir int

const (
	North Dir = iota
	East
	South
	West
)

func distance(input string, dir Dir) int {

	sum := 0
	coords := make(map[[2]int]byte)
	var pos [2]int

	for i, line := range strings.Split(input, "\n") {
		for j, c := range line {
			coords[[2]int{i, j}] = byte(c)
			if c == 'S' {
				pos = [2]int{i, j}
			}
		}
	}
	for {
		sum++
		switch dir {
		case North:
			pos[0]--
		case East:
			pos[1]++
		case South:
			pos[0]++
		case West:
			pos[1]--
		}

		if coords[pos] == '|' {
			continue
		} else if coords[pos] == '-' {
			continue
		} else if coords[pos] == '7' {
			if dir == East {
				dir = South
			} else {
				dir = West
			}
		} else if coords[pos] == 'F' {
			if dir == North {
				dir = East
			} else {
				dir = South
			}
		} else if coords[pos] == 'L' {
			if dir == South {
				dir = East
			} else {
				dir = North
			}
		} else if coords[pos] == 'J' {
			if dir == East {
				dir = North
			} else {
				dir = West
			}
		} else if coords[pos] == 'S' {
			break
		}
	}
	return sum / 2
}

func enclosed(input string, dir Dir) int {

	trail := make(map[[2]int]Dir)
	sum := 0
	coords := make(map[[2]int]byte)
	m := make([][]byte, 0)
	var pos [2]int

	for i, line := range strings.Split(input, "\n") {
		m = append(m, []byte(line))
		for j, c := range line {
			m[i][j] = byte(c)
			coords[[2]int{i, j}] = byte(c)
			if c == 'S' {
				pos = [2]int{i, j}
			}
		}
	}
	for {
		trail[pos] = dir

		switch dir {
		case North:
			pos[0]--
		case East:
			pos[1]++
		case South:
			pos[0]++
		case West:
			pos[1]--
		}

		if coords[pos] == '|' {
			continue
		} else if coords[pos] == '-' {
			continue
		} else if coords[pos] == '7' {
			if dir == East {
				dir = South
			} else {
				dir = West
			}
		} else if coords[pos] == 'F' {
			if dir == North {
				dir = East
			} else {
				dir = South
			}
		} else if coords[pos] == 'L' {
			if dir == South {
				dir = East
			} else {
				dir = North
			}
		} else if coords[pos] == 'J' {
			if dir == East {
				dir = North
			} else {
				dir = West
			}
		} else if coords[pos] == 'S' {
			break
		}

	}

	for i := 0; i < len(m); i++ {
		n := 0
		for j := 0; j < len(m[i]); j++ {
			c := m[i][j]
			if dir, ok := trail[[2]int{i, j}]; ok {
				if dir == South {
					n--
				} else {
					switch c {
					case '|':
						if dir == North {
							n++
						}
					case 'F':
						if dir == East {
							n++
						}
					case '7':
						if dir == West {
							n++
						}
					}
				}
			} else {
				if n != 0 {
					sum++
				}
			}
		}
	}

	return sum
}

func solve1(input string) int {
	distances := make([]int, 0)
	distances = append(distances, distance(input, South))
	distances = append(distances, distance(input, East))
	distances = append(distances, distance(input, West))
	distances = append(distances, distance(input, North))
	fmt.Println(distances)

	sort.IntSlice(distances).Sort()
	return distances[len(distances)-1]
}

func solve2(input string) int {
	// East is the longer path of MY input
	return enclosed(input, East)
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
