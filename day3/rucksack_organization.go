package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// TODO: if our set is less than 64 items, we can use a uint64 to represent it

type Item = rune
type Rucksack = string
type Set map[Item]struct{}

func main() {
	testInput := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	part1(strings.NewReader(testInput))

	b, _ := os.ReadFile("day3/input.txt")
	part1(bytes.NewReader(b))

	part2(strings.NewReader(testInput))
	part2(bytes.NewReader(b))
}

func part2(r io.Reader) {
	rucksacks := make(chan Rucksack)
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			rucksacks <- scanner.Text()
		}
		close(rucksacks)
	}()

	groups := make(chan [3]Rucksack)
	go func() {
		group := [3]Rucksack{}
		i := 0
		for rucksack := range rucksacks {
			group[i] = rucksack
			i++
			if i == 3 {
				groups <- group
				i = 0
				group = [3]Rucksack{}
			}
		}
		close(groups)
	}()

	priorities := make(chan int)
	go func() {
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
	}()

	total := 0
	for p := range priorities {
		total += p
	}

	fmt.Println(total)

}

func part1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		var leftSet, rightSet Set
		n := len(line)
		leftSet = newSet(line[:n/2])
		rightSet = newSet(line[n/2:])
		total += func() int {
			var ps []int
			for k, _ := range intersection(leftSet, rightSet) {
				ps = append(ps, priority(k))
			}
			if len(ps) != 1 {
				panic("The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.")
			}
			return ps[0]
		}()
	}

	fmt.Println(total)
}

func priority(i Item) int {
	var p int
	if 'A' <= i && i <= 'Z' {
		p = int(i-'A') + 27
	} else {
		p = int(i-'a') + 1
	}
	return p
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
