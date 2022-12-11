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

type Crate = string
type Stack = string

func main() {
	part1(strings.NewReader(testInput))
	part1(bytes.NewReader(b))

	part2(strings.NewReader(testInput))
	part2(bytes.NewReader(b))
}

func part1(r io.Reader) {
	solve(r, craneMover9000)
}

func part2(r io.Reader) {
	solve(r, craneMover9001)
}

func pop(s Stack) (Stack, Crate) {
	return s[1:], Crate(s[0])
}

func push(s Stack, c Crate) string {
	return c + s
}

func solve(r io.Reader, crane func(stacks []Stack, move, from, to int)) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()
	n := (len(line) + 1) / 4

	stacks := make([]Stack, n)
	for {
		if line[1] == '1' {
			break
		}

		for stackIndex := 0; stackIndex < n; stackIndex++ {
			crateIndex := (stackIndex * 4) + 1
			c := Crate(line[crateIndex])
			if c == " " {
				continue
			}
			stacks[stackIndex] = stacks[stackIndex] + c
		}

		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
	}

	// skip the empty line between stack state and rearrangement procedure
	scanner.Scan()

	for scanner.Scan() {
		var move, from, to int
		_, _ = fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &move, &from, &to)
		// cranes are 1 indexed in the rearrangement procedures
		from, to = from-1, to-1
		crane(stacks, move, from, to)
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Print(Crate(stacks[i][0]))
	}
	fmt.Println()
}

func craneMover9000(stacks []Stack, move, from, to int) {
	for i := 0; i < move; i++ {
		var crate string
		stacks[from], crate = pop(stacks[from])
		stacks[to] = push(stacks[to], crate)
	}
}

func craneMover9001(stacks []Stack, move, from, to int) {
	stacks[to] = stacks[from][:move] + stacks[to]
	stacks[from] = stacks[from][move:]
}
