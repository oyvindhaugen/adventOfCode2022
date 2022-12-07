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
	// for i := 0; i < len(stack); i++ {
	// 	for j := 0; j < len(stack[i]); j++ {
	// 		fmt.Printf("%s\n", stack[i][j])
	// 	}
	// 	fmt.Println()
	// }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//making the different params for moving crates
		str := scanner.Text()
		split := strings.Split(str, " ")
		qty, _ = strconv.Atoi(split[1])
		from, _ = strconv.Atoi(split[3])
		to, _ = strconv.Atoi(split[5])
		//moving the crates
		for i := 0; i < qty; i++ {
			crate, nStack := Pop(stack[from-1])
			stack[from-1] = nStack
			pStack := Push(stack[to-1], crate)
			stack[to-1] = pStack
		}
		// fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n", stack[0], stack[1], stack[2], stack[3], stack[4], stack[5], stack[6], stack[7], stack[8])
	}
	// fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", stack[0], stack[1], stack[2], stack[3], stack[4], stack[5], stack[6], stack[7], stack[8])
	for i := 0; i < len(stack); i++ {
		for j := 0; j < len(stack[i]); j++ {
			fmt.Printf("%s\n", stack[i][j])
		}
		fmt.Println()
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
func revSlice(slice []string) []string {
	reversed := []string{}
	for i := range slice {
		reversed = append(reversed, slice[len(slice)-1-i])
	}
	return reversed
}
