package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Guide:
	// Left column = opponents shape
	// Right column = our rune
	// A = Rock
	// B = Paper
	// C = Scissors

	// Scoring:
	// Total score = sum of score for each round
	// round score = shape score + outcome
	// shape score = 1 for Rock, 2 For Paper, 3 for Scissors
	// outcome score = 0 for loss, 3 for draw, 6 for win

	testEncryptedStrategyGuide := `A Y
B X
C Z`
	part1(strings.NewReader(testEncryptedStrategyGuide))

	b, _ := os.ReadFile("day2/input.txt")
	part1(bytes.NewReader(b))
	part2(bytes.NewReader(b))
}

type Shape = rune

const (
	Rock     = 'A'
	Paper    = 'B'
	Scissors = 'C'
)

var (
	shapeMap = map[rune]Shape{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
	}
)

func part1(r io.Reader) {
	scoreInput(r, recommendShape)
}

func part2(r io.Reader) {
	scoreInput(r, shapeByOutcome)
}

type ShapeChooser func(ours, theirs rune) Shape

func scoreInput(r io.Reader, chooseShape ShapeChooser) {
	scanner := bufio.NewScanner(r)

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		theirs := rune(line[0])
		ours := rune(line[2])
		score := scoreRound(chooseShape(ours, theirs), shapeMap[theirs])
		totalScore += score
	}

	fmt.Println(totalScore)
}

func scoreRound(ours, theirs Shape) int {
	type Score = int
	const (
		win  = 6
		draw = 3
		lose = 0
	)
	shapeScores := map[Shape]Score{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}
	matchupScores := map[Shape]map[Shape]Score{
		Rock:     {Rock: draw, Paper: lose, Scissors: win},
		Paper:    {Rock: win, Paper: draw, Scissors: lose},
		Scissors: {Rock: lose, Paper: win, Scissors: draw},
	}
	return shapeScores[ours] + matchupScores[ours][theirs]
}

func recommendShape(ours, _ rune) Shape {
	return map[rune]Shape{
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}[ours]
}

func shapeByOutcome(ours, theirs rune) Shape {
	type Outcome = rune
	const (
		lose = 'X'
		draw = 'Y'
		win  = 'Z'
	)
	outcomes := map[Outcome]map[Shape]Shape{
		win:  {Rock: Paper, Paper: Scissors, Scissors: Rock},
		lose: {Rock: Scissors, Paper: Rock, Scissors: Paper},
		draw: {Rock: Rock, Paper: Paper, Scissors: Scissors},
	}
	return outcomes[ours][theirs]
}
