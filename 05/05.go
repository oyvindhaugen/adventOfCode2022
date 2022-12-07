package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack := [9][]string{
		{"B", "Z", "T"},
		{"V", "H", "T", "D", "N"},
		{"B", "F", "M", "D"},
		{"T", "J", "G", "W", "V", "Q", "L"},
		{"W", "D", "G", "P", "V", "F", "Q", "M"},
		{"V", "Z", "Q", "G", "H", "F", "S"},
		{"Z", "S", "N", "R", "L", "T", "C", "W"},
		{"Z", "H", "W", "D", "J", "N", "R", "M"},
		{"M", "Q", "L", "F", "D", "S"},
	}
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var qty int
	var from int
	var to int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//making the different params for moving crates
		str := scanner.Text()
		split := strings.Split(str, " ")
		qty, _ = strconv.Atoi(split[1])
		from, _ = strconv.Atoi(split[3])
		to, _ = strconv.Atoi(split[5])
		for i := 0; i < qty; i++ {
			crate, nStack := Pop(stack[from-1])
			stack[from-1] = nStack
			pStack := Push(stack[to-1], crate)
			stack[to-1] = pStack
		}

	}
	for i := 0; i < len(stack); i++ {
		fmt.Printf("%s", stack[i][len(stack[i])-1])
	}
}
func Pop(stack []string) (string, []string) {
	e := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return e, stack
}
func Push(stack []string, item string) []string {
	stack = append(stack, item)
	return stack
}
