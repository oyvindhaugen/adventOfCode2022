package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func test(scanner *bufio.Scanner) {
	var layout [][]string
	line := 0
	for scanner.Scan() {
		height := strings.Split(scanner.Text(), "")
		var row []string
		for i := 0; i < len(height)-1; i++ {
			row = append(row, height[i])
		}
		layout = append(layout, row)
		line++
	}
	for _, num := range layout {
		fmt.Println(num)
	}
}

func main() {
	f, err := os.Open("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	test(scanner)
}
