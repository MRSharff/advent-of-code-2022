package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	part1(f)
}
