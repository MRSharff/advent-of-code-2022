package main

import (
	"fmt"
	"os"
)

var (
	testInput = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	b, _      = os.ReadFile("day6/input.txt")
)

func main() {
	part1(testInput)
	part1(string(b))

	part2(testInput)
	part2(string(b))
}

func part1(buffer string) {
	solve(buffer, 4)
}

func part2(buffer string) {
	solve(buffer, 14)
}

func solve(buffer string, markerSize int) {
	window := make(map[rune]int)

	hold := make(chan rune, markerSize)

	for i, c := range buffer {
		if i == markerSize {
			break
		}
		hold <- c
		window[c] = window[c] + 1
	}

	i := markerSize
	for _, entering := range buffer[i:] {
		if len(window) == markerSize {
			break
		}
		leaving := <-hold
		window[leaving]--
		if window[leaving] == 0 {
			delete(window, leaving)
		}
		hold <- entering
		window[entering] = window[entering] + 1
		i++
	}

	fmt.Println(i)
}
