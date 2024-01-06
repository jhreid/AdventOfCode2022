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

	results := count_visible(forest)
	fmt.Println(results[0])
	fmt.Println(results[1])
}

func count_visible(trees [][]int) []int {
	total := (len(trees) + len(trees[0]) - 2) * 2
	scenic_score := 0

	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[0])-1; col++ {
			current := trees[row][col]
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

			if current > max(west) || current > max(east) || current > max(north) || current > max(south) {
				total += 1
			}

			left := trees_in_view(current, reverse(west))
			right := trees_in_view(current, east)
			down := trees_in_view(current, south)
			up := trees_in_view(current, reverse(north))
			score := right * up * left * down
			if score > scenic_score {
				scenic_score = score
			}
		}
	}

	return []int{total, scenic_score}
}

func trees_in_view(tree int, s []int) int {
	in_view := 0
	if tree > max(s) {
		in_view = len(s)
	} else {
		for i, t := range s {
			if tree <= t {
				in_view = i + 1
				break
			}
		}
	}
	return in_view
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

func reverse(s []int) []int {
	r := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
