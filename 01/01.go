package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	firstPlace := 0
	secondPlace := 0
	thirdPlace := 0
	currentNumber := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			fmt.Println(scanner.Text())
			currNumb, _ := strconv.Atoi(scanner.Text())
			currentNumber += currNumb
		} else {
			first, second, third := placing(firstPlace, secondPlace, thirdPlace, currentNumber)
			firstPlace = first
			secondPlace = second
			thirdPlace = third
			currentNumber = 0
		}
	}
	fmt.Printf("Highest number: %d, %d, %d\n", firstPlace, secondPlace, thirdPlace)
	total := firstPlace + secondPlace + thirdPlace
	fmt.Println(total)
}
func placing(first int, second int, third int, current int) (int, int, int) {
	if current > first {
		third = second
		second = first
		first = current
		fmt.Printf("changed first: %d\n", first)
		return first, second, third
	}
	if current > second {
		third = second
		second = current
		fmt.Printf("changed second: %d\n", second)
		return first, second, third
	}
	if current > third {
		third = current
		fmt.Printf("changed third: %d\n", third)
		return first, second, third
	}
	return first, second, third
}
