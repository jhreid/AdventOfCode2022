package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	forest := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []int{}
		for _, r := range scanner.Text() {
			tree, err := strconv.Atoi(string(r))
			if err == nil {
				row = append(row, tree)
			}
		}
		forest = append(forest, row)
	}

	fmt.Println(count_visible(forest))
}

func count_visible(trees [][]int) int {
	total := (len(trees) + len(trees[0]) - 2) * 2

	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[0])-1; col++ {
			west := trees[row][:col]
			east := trees[row][col+1:]
			north := []int{}
			for _, t := range trees[:row] {
				north = append(north, t[col])
			}
			south := []int{}
			for _, t := range trees[row+1:] {
				south = append(south, t[col])
			}

			if trees[row][col] > max(west) || trees[row][col] > max(east) || trees[row][col] > max(north) || trees[row][col] > max(south) {
				total += 1
			}

		}
	}

	return total
}

func max(slice []int) int {
	max := slice[0]
	for _, m := range slice {
		if m > max {
			max = m
		}
	}
	return max
}
