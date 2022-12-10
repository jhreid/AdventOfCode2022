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

	x := 1
	cycles := 0
	values := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		switch input[0] {
		case "noop":
			{
				cycles += 1
				values = check_cycle(x, cycles, values)
			}
		case "addx":
			{
				cycles += 1
				values = check_cycle(x, cycles, values)
				cycles += 1
				values = check_cycle(x, cycles, values)
				val, _ := strconv.Atoi(input[1])
				x += val
			}
		}
	}

	fmt.Printf("x is %d\n", x)
	fmt.Printf("values %v\n", values)

	signal := sum(values)
	fmt.Printf("Total strength: %d\n", signal)
}

func check_cycle(x int, cycle int, values []int) []int {
	if (cycle+20)%40 == 0 {
		values = append(values, cycle*x)
	}
	return values
}

func sum(s []int) int {
	result := 0
	for _, v := range s {
		result += v
	}
	return result
}
