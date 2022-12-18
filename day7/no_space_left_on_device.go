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
}

func part1(r io.Reader) {
	root := NewDirectory(bufio.NewScanner(r))
	candidates := candidatesForDeletion(root)

	sum := 0
	for _, d := range candidates {
		sum += d.Size()
	}
	fmt.Println(sum)
}

func NewDirectory(scanner *bufio.Scanner) *Directory {
	d := &Directory{}
	var totalSize int
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "$ cd ..":
			return d
		case strings.HasPrefix(line, "dir") || line == "$ ls":
			continue
		case strings.HasPrefix(line, "$ cd"):
			nd := NewDirectory(scanner)
			totalSize += nd.size
			d.directories = append(d.directories, nd)
		default:
			var size int
			var fname string
			_, err := fmt.Sscanf(line, "%d %s", &size, &fname)
			if err != nil {
				panic(err)
			}
			totalSize += size
			d.files = append(d.files, &file{size: size})
		}
	}
	return d
}

func candidatesForDeletion(d *Directory) []*Directory {
	var candidates []*Directory

	if d.Size() <= 100000 {
		candidates = append(candidates, d)
	}
	for _, child := range d.directories {
		candidates = append(candidates, candidatesForDeletion(child)...)
	}
	return candidates
}

type file struct {
	size int
}

func (f *file) Size() int {
	return f.size
}

type Directory struct {
	name string
	size int

	directories []*Directory
	files       []*file
}

func (d *Directory) Size() int {
	total := 0
	for _, f := range d.files {
		total += f.Size()
	}
	for _, child := range d.directories {
		total += child.Size()
	}
	return total
}
