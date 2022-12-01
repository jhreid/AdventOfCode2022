package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	var totals []int
	var currentTotal int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totals = append(totals, currentTotal)
			currentTotal = 0
		} else {
			value, _ := strconv.Atoi(line)
			currentTotal += value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	fmt.Printf("Top amount: %v\n", totals[0])
	fmt.Printf("Top three total: %v", totals[0]+totals[1]+totals[2])
}
