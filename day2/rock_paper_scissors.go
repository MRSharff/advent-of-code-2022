package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Move int

const (
	Rock = iota
	Paper
	Scissors
)

var move = map[rune]Move{
	'A': Rock,
	'X': Rock,
	'B': Paper,
	'Y': Paper,
	'C': Scissors,
	'Z': Scissors,
}

var moveScore = map[Move]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func main() {
	testEncryptedStrategyGuide := `A Y
B X
C Z`

	// Guide:
	// Left column = opponents move
	// Right column = our move
	// A and X = Rock
	// B and Y = Paper
	// C and Z = Scissors

	// Scoring:
	// Total score = sum of score for each round
	// round score = shape score + outcome
	// shape score = 1 for Rock, 2 For Paper, 3 for Scissors
	// outcome score = 0 for loss, 3 for draw, 6 for win

	r := strings.NewReader(testEncryptedStrategyGuide)
	part1(r)

	b, _ := os.ReadFile("day2/input.txt")

	part1(bytes.NewReader(b))
}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		opponentMove := move[rune(line[0])]
		ourMove := move[rune(line[2])]

		score := moveScore[ourMove]
		if ourMove == opponentMove {
			score += 3
		} else if isWin(ourMove, opponentMove) {
			score += 6
		} // else we lose and our score does not improve
		totalScore += score
	}

	fmt.Println(totalScore)
}

func isWin(ourMove, opponentMove Move) bool {
	switch ourMove {
	case Rock: // our rock beats their scissors
		return opponentMove == Scissors
	case Paper: // our Paper beats their Rock
		return opponentMove == Rock
	case Scissors: // our Scissors beats their paper
		return opponentMove == Paper
	default:
		return false
	}
}
