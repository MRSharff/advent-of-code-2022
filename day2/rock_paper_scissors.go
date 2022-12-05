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

	scoreMap = map[Shape]int{
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

	shapeScore := map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}[shape[ours]]

	outcomeScore := 0

	if shape[ours] == theirs {
		outcomeScore = 3
	} else {
		win := false
		switch shape[ours] {
		case Rock: // our rock beats their scissors
			win = theirs == Scissors
		case Paper: // our Paper beats their Rock
			win = theirs == Rock
		case Scissors: // our Scissors beats their paper
			win = theirs == Paper
		}
		if win {
			outcomeScore = 6
		}
	} // else we lose and our score does not improve
	return shapeScore + outcomeScore
}

func outcomeChoosing(ours rune, theirs Shape) int {
	outcomes := map[rune]Outcome{
		'X': Lose,
		'Y': Draw,
		'Z': Win,
	}

	score := 0

	var shape Shape
	switch outcomes[ours] {
	case Win:
		shape = loseMap[theirs]
		score += 6
	case Draw:
		shape = theirs
		score += 3
	case Lose:
		shape = winMap[theirs]
		score += 0
	}
	score += scoreMap[shape]
	return score
}
