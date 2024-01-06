package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := string(file)
	parts := strings.Split(input, "\n\n")
	stacks := parts[0]
	instructions := strings.Split(parts[1], "\n")

	stack_map_one := process_stacks(stacks)
	stack_map_two := process_stacks(stacks)

	fmt.Println(part_one(instructions, stack_map_one))
	fmt.Println(part_two(instructions, stack_map_two))
}

func part_one(instructions []string, input map[int][]string) string {
	stacks := input
	output := ""

	for _, instr := range instructions {
		if instr == "" {
			break
		}
		tokens := strings.Split(instr, " ")
		n, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])

		for i := 0; i < n; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	for i := 1; i <= len(stacks); i++ {
		output = output + stacks[i][len(stacks[i])-1]
	}

	return output
}

func part_two(instructions []string, input map[int][]string) string {
	stacks := input
	output := ""

	for _, instr := range instructions {
		if instr == "" {
			break
		}
		tokens := strings.Split(instr, " ")
		n, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])

		for i := n; i > 0; i-- {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-i])
		}
		stacks[from] = stacks[from][:len(stacks[from])-n]
	}

	for i := 1; i <= len(stacks); i++ {
		output = output + stacks[i][len(stacks[i])-1]
	}

	return output
}

func process_stacks(stacks string) map[int][]string {
	lines := reverse(strings.Split(stacks, "\n"))
	stack_map := map[int][]string{}

	for i := 1; i <= 9; i++ {
		index := strings.Index(lines[0], fmt.Sprint(i))

		for l := 1; l < len(lines); l++ {
			char := lines[l][index]
			if char != ' ' {
				stack_map[i] = append(stack_map[i], string(char))
			}
		}
	}

	return stack_map
}

func reverse(slice []string) []string {
	var output []string

	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}

	return output
}
