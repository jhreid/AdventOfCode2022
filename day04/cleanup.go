package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		result := part_one(line)
		if result {
			total += 1
		}

		if result == false {
			fmt.Printf("%s %s\n", line, "false")
		}
	}

	fmt.Printf("Found %d overlaps\n", total)
}

func part_one(line string) bool {
	pairs := strings.Split(line, ",")

	a := strings.Split(pairs[0], "-")
	b := strings.Split(pairs[1], "-")

	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])

	if a0 <= b0 && a1 >= b1 {
		return true
	} else if b0 <= a0 && b1 >= a1 {
		return true
	}
	return false
}
