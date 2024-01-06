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

	fmt.Printf("Start of packet: %d\n", find_start(signal, 4))
	fmt.Printf("Start of message: %d\n", find_start(signal, 14))
}

func find_start(signal_buffer string, length int) int {
	found := 0
	for i := 0; i < len(signal_buffer)-length; i++ {
		if count_unique_chars(signal_buffer[i:i+length]) == length {
			found = i
			break
		}
	}

	return found + length
}

func count_unique_chars(s string) int {
	chars := make(map[rune]int)
	for _, c := range s {
		chars[c] = chars[c] + 1
	}
	return len(chars)
}
