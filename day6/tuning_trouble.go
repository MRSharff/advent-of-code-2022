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
}

func part1(buffer string) {
	window := make(map[rune]int)

	hold := make(chan rune, 4)

	for i, c := range buffer {
		if i == 4 {
			break
		}
		hold <- c
		window[c] = window[c] + 1
	}

	i := 4
	for _, entering := range buffer[i:] {
		if len(window) == 4 {
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
