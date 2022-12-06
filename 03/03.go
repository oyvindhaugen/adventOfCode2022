package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	totalPart1 := 0
	totalPart2 := 0

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dupe := split(scanner.Text())
		totalPart1 += findVal(dupe)

		if len(lines) == 2 {
			lines = append(lines, scanner.Text())
			sameChar := findChar(lines)
			totalPart2 += findVal(sameChar)
			lines = []string{}
		} else {
			lines = append(lines, scanner.Text())
		}
	}
	fmt.Println(totalPart1)
	fmt.Println(totalPart2)
}
func split(str string) (dupe string) {
	lengthOfStr := len(str)
	lengthOfStrDiv := lengthOfStr / 2
	comp1 := str[0:lengthOfStrDiv]
	comp2 := str[lengthOfStrDiv:lengthOfStr]
	for i := 0; i < lengthOfStrDiv; i++ {
		for j := 0; j < lengthOfStrDiv; j++ {
			if comp1[i] == comp2[j] {
				dupe = string(comp2[j])
				return
			}
		}
	}
	return
}
func findVal(char string) (val int) {
	f, err := os.Open("letters.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	i := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == char {
			val = i
			return val
		}
		i++
	}
	return val
}
func findChar(group []string) (char string) {
	first, second, third := group[0], group[1], group[2]
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			for k := 0; k < len(third); k++ {
				if first[i] == second[j] && second[j] == third[k] {
					char = string(third[k])
					fmt.Println(char)
					return char
				}
			}
		}
	}
	return char
}
