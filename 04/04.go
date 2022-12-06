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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totalContains := 0
	totalOverlaps := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sep := strings.Split(scanner.Text(), ",")
		contains := checkIfContains(sep)
		overlaps := checkIfOverlaps(sep)
		if contains {
			totalContains++
		}
		if overlaps {
			totalOverlaps++
		}
	}
	fmt.Println(totalContains)
	fmt.Println(totalOverlaps)
}
func checkIfContains(sep []string) (contains bool) {
	pair1 := strings.Split(sep[0], "-")
	pair2 := strings.Split(sep[1], "-")
	pair1First, _ := strconv.Atoi(pair1[0])
	pair1Second, _ := strconv.Atoi(pair1[1])
	pair2First, _ := strconv.Atoi(pair2[0])
	pair2Second, _ := strconv.Atoi(pair2[1])
	if pair2First >= pair1First && pair2Second <= pair1Second {
		contains = true
		return
	} else if pair1First >= pair2First && pair1Second <= pair2Second {
		contains = true
		return
	}
	return
}
func checkIfOverlaps(sep []string) (overlaps bool) {
	pair1 := strings.Split(sep[0], "-")
	pair2 := strings.Split(sep[1], "-")
	oneOne, _ := strconv.Atoi(pair1[0])
	oneTwo, _ := strconv.Atoi(pair1[1])
	twoOne, _ := strconv.Atoi(pair2[0])
	twoTwo, _ := strconv.Atoi(pair2[1])
	if oneOne >= twoOne && oneTwo <= twoTwo || oneTwo >= twoOne && oneTwo <= twoTwo || twoOne >= oneOne && twoOne <= oneTwo || twoTwo >= oneOne && twoTwo <= oneTwo {
		overlaps = true
		return
	}
	return
}
