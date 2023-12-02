package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	total  int = 70000000
	needed int = 30000000
)

func main() {
	terminalLog := readTerminalLog("day7/input.txt")
	treeScanner := newTreeScanner()
	for _, ln := range terminalLog {
		treeScanner.scan(ln)
	}

	used := treeScanner.root.size()
	free := total - used
	threshold := needed - free
	fmt.Printf("Used: %v, Free: %v, Threshold: %v\n", used, free, threshold)
	if threshold < 0 {
		panic("Threshold is negative")
	}

	smallestOverThreshold := 0
	dirs := treeScanner.root.flatList()
	for _, d := range dirs {
		size := d.size()
		if size >= threshold &&
			(smallestOverThreshold == 0 || size < smallestOverThreshold) {
			fmt.Printf("New smallest: %v %v\n", d.name, size)
			smallestOverThreshold = size
		}
	}
	fmt.Printf("Smallest over %v: %v\n", threshold, smallestOverThreshold)
}

type TreeScanner struct {
	readingLS bool
	root      *Directory
	curDir    *Directory
}

func newTreeScanner() *TreeScanner {
	root := &Directory{
		name: "/",
	}
	return &TreeScanner{
		readingLS: false,
		root:      root,
		curDir:    root,
	}
}

func (t *TreeScanner) scan(ln string) {
	if strings.HasPrefix(ln, "$") {
		// It's a new command
		command := strings.Fields(ln)[1]
		switch command {
		case "ls":
			t.readingLS = true
		case "cd":
			t.readingLS = false
			t.cd(strings.Fields(ln)[2])
		}
	} else {
		if !t.readingLS {
			panic("No command and not reading ls output")
		}
		if strings.HasPrefix(ln, "dir") {
			// Line is a directory
			dir := &Directory{
				name:   strings.TrimSpace(strings.TrimPrefix(ln, "dir")),
				parent: t.curDir,
			}
			t.curDir.children = append(t.curDir.children, dir)
		} else {
			// Line is a file
			size, err := strconv.Atoi(strings.Fields(ln)[0])
			if err != nil {
				panic(err)
			}
			file := &File{
				filesize: size,
			}
			t.curDir.files = append(t.curDir.files, file)
		}
	}
}

func (t *TreeScanner) cd(dirname string) {
	switch dirname {
	case "/":
		t.curDir = t.root
	case "..":
		t.curDir = t.curDir.parent
	default:
		// Check if dirname is an existing child of t.curDir
		for _, c := range t.curDir.children {
			if c.name == dirname {
				t.curDir = c
				return
			}
		}
		// It's not, create it and make the change
		dir := &Directory{
			name:   dirname,
			parent: t.curDir,
		}
		t.curDir.children = append(t.curDir.children, dir)
		t.curDir = dir
	}
}

type Directory struct {
	name     string
	parent   *Directory
	children []*Directory
	files    []*File
}

func (d *Directory) path() string {
	if d.parent == nil {
		return "/"
	}
	return d.parent.path() + d.name + "/"
}

func (d *Directory) size() int {
	total := 0
	for _, c := range d.children {
		total += c.size()
	}
	for _, f := range d.files {
		total += f.filesize
	}
	return total
}

func (d *Directory) flatList() []*Directory {
	result := []*Directory{d}
	for _, c := range d.children {
		result = append(result, c.flatList()...)
	}
	return result
}

type File struct {
	filesize int
}

func readTerminalLog(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		ln := scanner.Text()
		result = append(result, ln)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
