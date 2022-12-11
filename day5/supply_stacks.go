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
	testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	b, _ = os.ReadFile("day5/input.txt")
)

func main() {
	part1(strings.NewReader(testInput))
	part1(bytes.NewReader(b))
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()
	n := (len(line) + 1) / 4

	stacks := make([][]string, n)
	for {
		if line[1] == '1' {
			break
		}

		for i, x := 1, 0; i < len(line); i, x = i+4, x+1 {
			k := string(line[i])
			if k == " " {
				continue
			}
			stacks[x] = append(stacks[x], k)
		}

		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
	}

	for scanner.Scan() {
		var count, from, to int
		_, _ = fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to)
		for i := 0; i < count; i++ {
			var crate string
			stacks[from-1], crate = pop(stacks[from-1])
			stacks[to-1] = push(stacks[to-1], crate)
		}
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Print(stacks[i][0])
	}
	fmt.Println()
}

func pop(stack []string) ([]string, string) {
	return stack[1:], stack[0]
}

func push(stack []string, s string) []string {
	return append([]string{s}, stack...)
}
