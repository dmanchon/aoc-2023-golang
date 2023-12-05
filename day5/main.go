package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Source      int64
	Destination int64
	Lenght      int64
}

func solve1(input string) int64 {
	maps := make([][]Rule, 0)
	sources := make([]int64, 0)
	for i, line := range strings.Split(input, "\n\n") {
		m := make([]Rule, 0)
		split1 := strings.Split(line, "\n")

		for _, s := range split1[1:] {
			split2 := strings.Split(s, " ")
			if i == 0 {
				for _, n := range split2 {
					src, _ := strconv.ParseInt(n, 10, 64)
					sources = append(sources, src)
				}
				break
			} else {
				dst, _ := strconv.ParseInt(split2[0], 10, 64)
				src, _ := strconv.ParseInt(split2[1], 10, 64)
				lenght, _ := strconv.ParseInt(split2[2], 10, 64)
				n := Rule{src, dst, lenght}
				m = append(m, n)
			}
		}
		maps = append(maps, m)
	}

	for _, m := range maps {
		for i, source := range sources {
			for _, target := range m {
				if source >= target.Source && source < target.Source+target.Lenght {
					sources[i] = (target.Destination - target.Source) + source
					break
				}
			}
		}
	}
	// get minimun value of sources
	min := sources[0]
	for _, source := range sources {
		if source < min {
			min = source
		}
	}
	return min
}

func solve2(input string) int64 {
	maps := make([][]Rule, 0)
	sources := make([]int64, 0)
	for i, line := range strings.Split(input, "\n\n") {
		m := make([]Rule, 0)
		split1 := strings.Split(line, "\n")

		for _, s := range split1[1:] {
			split2 := strings.Split(s, " ")
			if i == 0 {
				for i := 0; i < len(split2); i += 2 {
					src, _ := strconv.ParseInt(split2[i], 10, 64)
					l, _ := strconv.ParseInt(split2[i+1], 10, 64)
					// lets use brute force
					for j := int64(0); j < l; j++ {
						sources = append(sources, src+j)
					}
				}
				break
			} else {
				dst, _ := strconv.ParseInt(split2[0], 10, 64)
				src, _ := strconv.ParseInt(split2[1], 10, 64)
				lenght, _ := strconv.ParseInt(split2[2], 10, 64)
				n := Rule{src, dst, lenght}
				m = append(m, n)
			}
		}
		maps = append(maps, m)
	}

	for _, m := range maps {
		for i, source := range sources {
			for _, target := range m {
				if source >= target.Source && source < target.Source+target.Lenght {
					sources[i] = (target.Destination - target.Source) + source
					break
				}
			}
		}
	}
	// get minimun value of sources
	min := sources[0]
	for _, source := range sources {
		if source < min {
			min = source
		}
	}

	return min
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
