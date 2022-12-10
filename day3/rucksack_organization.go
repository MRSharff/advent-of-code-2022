package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/MRSharff/advent-of-code-2022/metrics"
)

var (
	testInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	b, _ = os.ReadFile("day3/input.txt")
)

type Item = rune
type Rucksack = string
type Set map[Item]struct{}

func main() {
	part1(strings.NewReader(testInput))

	part1(bytes.NewReader(b))

	part2(strings.NewReader(testInput))
	metrics.Timeit(func() { part2(bytes.NewReader(b)) })
}

func part2(r io.Reader) {
	solve(r, byGroup)
}

func part1(r io.Reader) {
	solve(r, byRucksack)
}

type Prioritizer func(rucksacks chan Rucksack, priorities chan int)

func solve(r io.Reader, prioritize Prioritizer) {
	rucksacks := make(chan Rucksack, 300)
	go readRucksacks(r, rucksacks)

	priorities := make(chan int, 300)
	go prioritize(rucksacks, priorities)

	total := 0
	for p := range priorities {
		total += p
	}

	fmt.Println(total)
}

func byRucksack(rucksacks chan Rucksack, priorities chan int) {
	for rucksack := range rucksacks {
		var leftSet, rightSet Set
		n := len(rucksack)
		leftSet = newSet(rucksack[:n/2])
		rightSet = newSet(rucksack[n/2:])
		itemsInBothCompartments := intersection(leftSet, rightSet)
		if len(itemsInBothCompartments) != 1 {
			panic("The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.")
		}
		for k, _ := range itemsInBothCompartments {
			priorities <- priority(k)
		}
	}
	close(priorities)
}

func byGroup(rucksacks chan Rucksack, priorities chan int) {
	groups := make(chan [3]Rucksack, 100)
	go readGroups(rucksacks, groups)

	for group := range groups {
		itemsCarriedByAll3 := intersection(newSet(group[0]), intersection(newSet(group[1]), newSet(group[2])))
		if len(itemsCarriedByAll3) != 1 {
			panic("the badge is the only item type carried by all three Elves")
		}
		for k, _ := range itemsCarriedByAll3 {
			priorities <- priority(k)
		}
	}
	close(priorities)
}

func readRucksacks(r io.Reader, rucksacks chan Rucksack) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rucksacks <- scanner.Text()
	}
	close(rucksacks)
}

func readGroups(rucksacks chan Rucksack, groups chan [3]Rucksack) {
	for rucksack := range rucksacks {
		groups <- [3]Rucksack{rucksack, <-rucksacks, <-rucksacks}
	}
	close(groups)
}

func priority(i Item) int {
	if 'A' <= i && i <= 'Z' {
		return int(i-'A') + 27
	}
	return int(i-'a') + 1

}

func newSet(runes string) Set {
	set := Set{}
	for _, c := range runes {
		set[c] = struct{}{}
	}
	return set
}

func intersection(a, b Set) Set {
	set := Set{}
	for k, _ := range a {
		if _, ok := b[k]; ok {
			set[k] = struct{}{}
		}
	}
	return set
}
