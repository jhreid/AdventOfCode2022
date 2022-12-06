package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	signal := string(file)

	fmt.Printf("Start: %d\n", find_start_of_packet(signal))
}

func find_start_of_packet(signal_buffer string) int {
	found := 0
	for i := 0; i < len(signal_buffer); i++ {
		if count_unique_chars(signal_buffer[i:i+4]) == 4 {
			found = i
			break
		}
	}

	return found + 4
}

func count_unique_chars(s string) int {
	fmt.Println(s)
	chars := make(map[rune]int)
	for _, c := range s {
		chars[c] = chars[c] + 1
	}
	return len(chars)
}
