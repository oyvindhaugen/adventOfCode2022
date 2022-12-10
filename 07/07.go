package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	which    int // 0: Directory, 1: File
	size     int
	parent   *File
	children map[string]*File
}

func createDirectory(parentFile *File) *File {
	return &File{
		which:    0,
		parent:   parentFile,
		children: map[string]*File{},
	}
}

func dirSize(scanner *bufio.Scanner) {
	files := []*File{}
	indent := map[*File]int{}
	rootDir := createDirectory(nil)
	rootDir.parent = rootDir
	node := rootDir

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		switch fields[0] {
		case "$":
			switch fields[1] {
			case "cd":
				switch fields[2] {
				case "/":
					node = rootDir
				case "..":
					node = node.parent
				default:
					node = node.children[fields[2]]
				}
			}
		case "dir":
			node.children[fields[1]] = createDirectory(node)
			indent[node] += 1
		default:
			node.children[fields[1]] = &File{
				which:  1,
				size:   toInt(fields[0]),
				parent: node,
			}
			indent[node] += 1
			files = append(files, node.children[fields[1]])
		}
	}
	sizes := []int{0}
	for i := 0; i < len(files); i++ {
		file := files[i]
		parent := file.parent
		if parent != file {
			parent.size += file.size
			indent[parent] -= 1
			if indent[parent] == 0 {
				files = append(files, parent)
			}
		}
		if file.which == 0 {
			sizes = append(sizes, file.size)
		}
	}

	sort.Ints(sizes)
	remaining := 70000000 - sizes[len(sizes)-1]
	required := 30000000
	total := 0
	for _, size := range sizes {
		if size+remaining >= required {
			fmt.Println("Part 2: ", size)
			return
		} else if size < 100001 {
			total += size
		}
	}
	fmt.Println("Part 1: ", total)
}

func toInt(num string) int {
	v, _ := strconv.ParseInt(num, 10, 64)
	return int(v)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	dirSize(scanner)
}

/*
read the current line
do a switch on what statement (cd or ls)
if ls then append all the filesizes into a slice
add the slice together and check if it is 100,000 or less
if true then add it to the total
*/
