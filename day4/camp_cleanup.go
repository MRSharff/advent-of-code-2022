package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	testInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	b, _ = os.ReadFile("day4/input.txt")
)

func main() {
	part1(strings.NewReader(testInput))
	part1(bytes.NewReader(b))
}

func part1(r io.Reader) {
	var overlappers []string
	for line := range lines(r) {
		var a, b, x, y int
		_, _ = fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &x, &y)
		if fullyContains(a, b, x, y) {
			overlappers = append(overlappers, line)
		}

	}
	fmt.Println(len(overlappers))
}

func fullyContains(a, b, x, y int) bool {
	return (a <= x && y <= b) || (x <= a && b <= y)
}

func lines(r io.Reader) chan string {
	l := make(chan string)
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			l <- scanner.Text()
		}
		close(l)
	}()
	return l
}
