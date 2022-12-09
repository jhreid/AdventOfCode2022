package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	locations := make(map[Pos]int)
	tails := []Pos{}
	head := Pos{0, 0}
	tail := Pos{0, 0}
	locations[tail] = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			instruction := strings.Split(line, " ")
			direction := instruction[0]
			distance, _ := strconv.Atoi(instruction[1])

			for d := 1; d <= distance; d++ {
				tx := tail.x
				ty := tail.y
				switch direction {
				case "R":
					{
						head = Pos{head.x + 1, head.y}
						if head.x-tx > 1 {
							tx += 1
							if ty != head.y && d > 1 {
								ty = head.y
							}
						}
					}
				case "U":
					{
						head = Pos{head.x, head.y + 1}
						if head.y-ty > 1 {
							ty += 1
							if tx != head.x && d > 1 {
								tx = head.x
							}
						}
					}
				case "L":
					{
						head = Pos{head.x - 1, head.y}
						if head.x-tx < -1 {
							tx -= 1
							if ty != head.y && d > 1 {
								ty = head.y
							}
						}
					}
				case "D":
					{
						head = Pos{head.x, head.y - 1}
						if head.y-ty < -1 {
							ty -= 1
							if tx != head.x && d > 1 {
								tx = head.x
							}
						}
					}
				}
				tail = Pos{tx, ty}
				locations[tail] += 1
				tails = append(tails, tail)
			}

			// fmt.Printf("Head position: %v\n", head)
			// fmt.Printf("Tail position: %v\n", tail)
		}
	}

	fmt.Printf("Positions visited: %d\n", len(locations))
}
