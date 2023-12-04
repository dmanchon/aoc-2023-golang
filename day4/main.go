package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Card struct {
	id      int
	numbers map[int]interface{}
	winners map[int]interface{}
}

func ParseLine(line string) Card {
	split1 := strings.Split(line, ":")
	split2 := strings.Split(split1[1], "|")
	split3 := strings.Split(split1[0], " ")
	id := 0
	fmt.Sscanf(split3[len(split3)-1], "%d", &id)

	winners := make(map[int]interface{})
	numbers := make(map[int]interface{})
	for _, num_str := range strings.Split(split2[0], " ") {
		if num_str == "" {
			continue
		}
		num := 0
		fmt.Sscanf(num_str, "%d", &num)
		numbers[num] = struct{}{}
	}

	for _, num_str := range strings.Split(split2[1], " ") {
		if num_str == "" {
			continue
		}
		num := 0
		fmt.Sscanf(num_str, "%d", &num)
		winners[num] = struct{}{}
	}
	return Card{id, numbers, winners}
}
func Points(card Card) int {
	points := 0
	for num := range card.numbers {
		if _, ok := card.winners[num]; ok {
			points += 1
		}
	}
	return points

}

func solve1(input string) int {
	sum := 0.0
	for _, line := range strings.Split(input, "\n") {
		card := ParseLine(line)

		winning := Points(card)
		if winning >= 1 {
			sum += math.Pow(2, float64(winning)-1.0)
		}
	}
	return int(sum)
}

func solve2(input string) int {
	counts := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		card := ParseLine(line)
		counts[card.id] += 1
		winning := Points(card)
		for i := card.id + 1; i < card.id+1+winning; i++ {
			counts[i] += counts[card.id]
		}
	}
	sum := 0
	for _, v := range counts {
		sum += v
	}
	return sum
}

func main() {
	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	fmt.Printf("Part1: %d\nPart2: %d\n", solve1(input), solve2(input))
}
