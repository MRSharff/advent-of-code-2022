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

type Shape int
type Outcome int

const (
	Rock Shape = iota
	Paper
	Scissors

	Win Outcome = iota
	Lose
	Draw
)

var (
	shapeMap = map[rune]Shape{
		'A': Rock, 'X': Rock,
		'B': Paper, 'Y': Paper,
		'C': Scissors, 'Z': Scissors,
	}

	shapeScores = map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}

	winMap = map[Shape]Shape{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}

	loseMap = map[Shape]Shape{
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	}
)

func part1(r io.Reader) {
	scoreInput(r, shapeChoosing)
}

func part2(r io.Reader) {
	scoreInput(r, outcomeChoosing)
}

type Scorer func(ours rune, theirs Shape) int

func scoreInput(r io.Reader, scorer Scorer) {
	scanner := bufio.NewScanner(r)

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		theirs := rune(line[0])
		ours := rune(line[2])
		score := scorer(ours, shapeMap[theirs])
		totalScore += score
	}

	fmt.Println(totalScore)
}

func shapeChoosing(ours rune, theirs Shape) int {
	var shape = map[rune]Shape{
		'A': Rock, 'X': Rock,
		'B': Paper, 'Y': Paper,
		'C': Scissors, 'Z': Scissors,
	}

	return scoreRound(shape[ours], theirs)
}

func outcomeChoosing(ours rune, theirs Shape) int {
	outcomes := map[rune]Outcome{
		'X': Lose,
		'Y': Draw,
		'Z': Win,
	}

	var shape Shape
	switch outcomes[ours] {
	case Win:
		shape = loseMap[theirs]
	case Draw:
		shape = theirs
	case Lose:
		shape = winMap[theirs]
	}
	return scoreRound(shape, theirs)
}

func scoreRound(ours, theirs Shape) int {
	const (
		win  = 6
		draw = 3
		lose = 0
	)
	outcomeScores := map[Shape]map[Shape]int{
		Rock:     {Rock: draw, Paper: lose, Scissors: win},
		Paper:    {Rock: win, Paper: draw, Scissors: lose},
		Scissors: {Rock: lose, Paper: win, Scissors: draw},
	}
	return shapeScores[ours] + outcomeScores[ours][theirs]
}
