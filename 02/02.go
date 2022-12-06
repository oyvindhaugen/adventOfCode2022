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

	score := 0
	score2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		score += part1(scanner.Text())
		score2 += part2(scanner.Text())
	}
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Score2: %d\n", score2)
}
func part1(fight string) (score int) {
	switch fight {
	case "A X":
		score += 4
	case "A Y":
		score += 8
	case "A Z":
		score += 3
	case "B X":
		score += 1
	case "B Y":
		score += 5
	case "B Z":
		score += 9
	case "C X":
		score += 7
	case "C Y":
		score += 2
	case "C Z":
		score += 6
	}
	return
}

// x = lose
// y = draw
// z = win
func part2(fight string) (score int) {
	switch fight {
	case "A X":
		score += 3
	case "A Y":
		score += 4
	case "A Z":
		score += 8
	case "B X":
		score += 1
	case "B Y":
		score += 5
	case "B Z":
		score += 9
	case "C X":
		score += 2
	case "C Y":
		score += 6
	case "C Z":
		score += 7
	}
	return
}
