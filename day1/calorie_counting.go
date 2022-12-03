package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func part1(input io.Reader) {
	max := 0
	total := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if total > max {
				max = total
			}
			total = 0
			continue
		}
		var calories int
		_, _ = fmt.Sscanf(line, "%d", &calories)
		total += calories
	}
	fmt.Println(max)
}

func part2(input io.Reader) {
	var calories []int
	total := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calories = append(calories, total)
			total = 0
			continue
		}
		var calorie int
		_, _ = fmt.Sscanf(line, "%d", &calorie)
		total += calorie
	}

	sort.Ints(calories)
	total = 0
	for i := len(calories) - 1; i >= len(calories)-3; i-- {
		total += calories[i]
	}
	fmt.Println(total)
}

func main() {
	// input is a list of calorie counts
	// one item per line
	// one elf per \n\n

	testInput := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

	r := strings.NewReader(testInput)
	part1(r)

	f, _ := os.Open("day1/input.txt")
	defer f.Close()

	input, _ := io.ReadAll(f)

	part1(bytes.NewReader(input))
	part2(bytes.NewReader(input))
}
