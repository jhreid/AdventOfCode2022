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

	total1 := 0
	total2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		result1 := part_one(line)
		if result1 {
			total1 += 1
		}

		result2 := part_two(line)
		if result2 {
			total2 += 1
		}
	}

	fmt.Printf("Found %d fully contained\n", total1)
	fmt.Printf("Found %d foverlaps\n", total2)
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

func part_two(line string) bool {
	pairs := strings.Split(line, ",")

	a := strings.Split(pairs[0], "-")
	b := strings.Split(pairs[1], "-")

	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])

	if a0 >= b0 && a0 <= b1 {
		return true
	} else if b0 >= a0 && b0 <= a1 {
		return true
	}

	return false
}
