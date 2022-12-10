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

	part2(strings.NewReader(testInput))
	part2(bytes.NewReader(b))
}

func solve(r io.Reader, overlapFunc func(a, b, x, y int) bool) {
	scanner := bufio.NewScanner(r)
	var overlapCount int
	for scanner.Scan() {
		var a, b, x, y int
		_, _ = fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &x, &y)
		if overlapFunc(a, b, x, y) {
			overlapCount++
		}
	}
	fmt.Println(overlapCount)
}

func part1(r io.Reader) {
	solve(r, fullyContains)
}

func fullyContains(a, b, x, y int) bool {
	return (a <= x && y <= b) || (x <= a && b <= y)
}

func part2(r io.Reader) {
	solve(r, anyOverlap)
}

func anyOverlap(a, b, x, y int) bool {
	return (a <= x && x <= b) || (x <= a && a <= y)
}
