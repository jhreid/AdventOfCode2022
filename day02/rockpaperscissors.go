package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	var scores = make(map[string]int)
	scores["A X"] = 4
	scores["A Y"] = 8
	scores["A Z"] = 3
	scores["B X"] = 1
	scores["B Y"] = 5
	scores["B Z"] = 9
	scores["C X"] = 7
	scores["C Y"] = 2
	scores["C Z"] = 6

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		score += scores[scanner.Text()]
	}

	fmt.Printf("Score: %d\n", score)
}

// A X = 4
// A Y = 8
// A Z = 3

// B X = 1
// B Y = 5
// B Z = 9

// C X = 7
// C Y = 2
// C Z = 6
