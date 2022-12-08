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
	startMarker := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		startMarker = append(startMarker, string(scanner.Text()[0]))
		startMarker = append(startMarker, string(scanner.Text()[1]))
		startMarker = append(startMarker, string(scanner.Text()[2]))
		for i := 3; i < len(scanner.Text()); i++ {
			doesContain, where := contains(startMarker, string(scanner.Text()[i]))
			if doesContain {
				startMarker = startMarker[where:]
				startMarker = append(startMarker, string(scanner.Text()[i]))
			} else if len(startMarker) == 14 {
				fmt.Println(i)
				break
			} else {
				startMarker = append(startMarker, string(scanner.Text()[i]))
			}
		}
	}
}
func contains(slice []string, elem string) (doesContain bool, which int) {
	for i := 0; i < len(slice); i++ {
		if elem == slice[i] {
			doesContain = true
			which = i + 1
		}
	}
	return
}

type void struct{}

var member void
