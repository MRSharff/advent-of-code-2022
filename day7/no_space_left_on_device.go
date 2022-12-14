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
	testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	b, _ = os.ReadFile("day7/input.txt")
)

func main() {
	part1(strings.NewReader(testInput))
	part1(bytes.NewReader(b))

	part2(strings.NewReader(testInput))
	part2(bytes.NewReader(b))
}

func part1(r io.Reader) {
	root := NewDirectory(bufio.NewScanner(r))

	sum := 0
	Visit(root, func(d *Directory) {
		if d.Size() <= 100_000 {
			sum += d.Size()
		}
	})

	fmt.Println(sum)
}

func part2(r io.Reader) {
	root := NewDirectory(bufio.NewScanner(r))

	totalDiskSpace := 70_000_000
	totalRequiredSpace := 30_000_000

	totalDiskSpaceUsed := root.Size()
	unusedDiskSpace := totalDiskSpace - totalDiskSpaceUsed
	requiredSpace := totalRequiredSpace - unusedDiskSpace

	var smallestDirGreaterThanRequired = root

	Visit(smallestDirGreaterThanRequired, func(d *Directory) {
		if d.Size() <= smallestDirGreaterThanRequired.Size() && d.Size() >= requiredSpace {
			smallestDirGreaterThanRequired = d
		}
	})

	fmt.Println(smallestDirGreaterThanRequired.Size())
}

func NewDirectory(scanner *bufio.Scanner) *Directory {
	d := &Directory{}
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "$ cd ..":
			return d
		case strings.HasPrefix(line, "dir") || line == "$ ls":
			continue
		case strings.HasPrefix(line, "$ cd"):
			nd := NewDirectory(scanner)
			d.size += nd.size
			d.directories = append(d.directories, nd)
		default:
			var size int
			_, err := fmt.Sscanf(line, "%d ", &size)
			if err != nil {
				panic(err)
			}
			d.size += size
		}
	}
	return d
}

type Directory struct {
	name string
	size int

	directories []*Directory
}

func (d *Directory) Size() int {
	return d.size
}

func Visit(d *Directory, visitorFunc func(*Directory)) {
	visitorFunc(d)
	for _, d := range d.directories {
		Visit(d, visitorFunc)
	}
}
