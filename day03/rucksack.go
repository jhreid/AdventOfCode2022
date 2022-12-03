package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	duplicates := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		contents := strings.Split(line, "")
		compartmentone := contents[:len(contents)/2]
		compartmenttwo := contents[len(contents)/2:]

		// fmt.Println(compartmentone)
		// fmt.Println(compartmenttwo)

		for _, item := range compartmentone {
			if slices.Contains(compartmenttwo, item) {
				duplicates = append(duplicates, item)
				break
			}
		}
	}

	fmt.Println(duplicates)

	score := 0
	values_map := get_values_map()
	for _, item := range duplicates {
		score += values_map[item]
	}

	fmt.Printf("Total score: %d\n", score)
}

func get_values_map() map[string]int {
	values_map := map[string]int{}

	score := 0
	for i := 'a'; i <= 'z'; i++ {
		score += 1
		values_map[string(i)] = score
	}
	for i := 'A'; i <= 'Z'; i++ {
		score += 1
		values_map[string(i)] = score
	}

	return values_map
}
