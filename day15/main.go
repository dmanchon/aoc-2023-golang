package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Hash(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash = hash % 256
	}
	fmt.Print(s, " ", hash, "\n")
	return hash

}
func solve1(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		for _, c := range strings.Split(line, ",") {
			sum += Hash(c)
		}
	}
	return sum
}

type Lens struct {
	Value string
	Len   int
}

func solve2(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		m := make(map[int][]Lens)
		for _, c := range strings.Split(line, ",") {
			if strings.Contains(c, "=") {
				b := strings.Split(c, "=")[0]
				ns := strings.Split(c, "=")[1]
				n, _ := strconv.Atoi(ns)
				hash := Hash(b)
				l := Lens{Value: b, Len: n}
				if lens, ok := m[hash]; !ok {
					m[hash] = make([]Lens, 0)
					lens = append(lens, l)
					m[hash] = lens
				} else {
					// replace or append l to the list
					ok := false
					for i, l2 := range lens {
						if l2.Value == b {
							lens[i] = l
							ok = true
						}
					}
					if !ok {
						lens = append(lens, l)
					}
					m[hash] = lens
				}
			} else if strings.Contains(c, "-") {
				b := strings.Split(c, "-")
				hash := Hash(b[0])
				if lens, ok := m[hash]; ok {
					// remove l from lens
					for i, l := range lens {
						if l.Value == b[0] {
							lens = append(lens[:i], lens[i+1:]...)
						}
					}
					m[hash] = lens
				}
			}
		}
		for box, lens := range m {
			for j, l := range lens {
				n := (box + 1) * (j + 1) * l.Len
				sum += n
			}
		}
	}
	return sum
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
