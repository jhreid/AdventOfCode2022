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

	var scores_2 = make(map[string]int)
	scores_2["A X"] = 3 // lose 0 + 3
	scores_2["A Y"] = 4 // draw 3 + 1
	scores_2["A Z"] = 8 // win 6 + 2
	scores_2["B X"] = 1 // lose 0 + 1
	scores_2["B Y"] = 5 // draw 3 + 2
	scores_2["B Z"] = 9 // win 6 + 3
	scores_2["C X"] = 2 // lose 0 + 2
	scores_2["C Y"] = 6 // draw 3 + 3
	scores_2["C Z"] = 7 // win 6 + 1

	score := 0
	score_2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		score += scores[line]
		score_2 += scores_2[line]
	}

	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Score part 2: %d\n", score_2)
}
