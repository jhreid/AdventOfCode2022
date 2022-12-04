package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	badges := []string{}
	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		a := strings.Split(line, "")
		line, _ = reader.ReadString('\n')
		b := strings.Split(line, "")
		line, _ = reader.ReadString('\n')
		c := strings.Split(line, "")

		badge := ""
		for _, item := range a {
			if slices.Contains(b, item) && slices.Contains(c, item) {
				badge = item
				break
			}
		}
		fmt.Printf("found badge %s\n", badge)

		badges = append(badges, badge)
	}

	fmt.Println(badges)

	score := 0
	values_map := get_values_map()
	for _, item := range badges {
		score += values_map[item]
	}

	fmt.Printf("Total score: %d\n", score)
}

func get_values_map() map[string]int {
	values_map := map[string]int{}

	score := 0
	for i := 'a'; i <= 'z'; i++ {
		score += 1
		values_map[string(i)] = score
	}
	for i := 'A'; i <= 'Z'; i++ {
		score += 1
		values_map[string(i)] = score
	}

	return values_map
}
