package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`^Valve (\w{2}) .* rate=(\d*); .* to valve[s]* (.*)`)
	for scanner.Scan() {
		line := scanner.Text()
		found := re.FindStringSubmatch(line)
	}
}
