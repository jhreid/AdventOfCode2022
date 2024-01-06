package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

func makePos(s string) Pos {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Pos{x, y}
}

type Cave map[Pos]int

func (c Cave) maxY() int {
	y := 0
	for p, _ := range c {
		if p.y > y {
			y = p.y
		}
	}
	return y
}

func (c Cave) countSand() int {
	s := 0
	for _, v := range c {
		if v == 2 {
			s++
		}
	}
	return s
}

func (c Cave) addRocks(input string) {
	points := strings.Split(input, " -> ")
	vertices := []Pos{}
	for _, p := range points {
		vertices = append(vertices, makePos(p))
	}

	for i := 1; i < len(vertices); i++ {
		from := vertices[i-1]
		to := vertices[i]

		c[from] = 1
		c[to] = 1

		if from.x < to.x {
			for i := from.x + 1; i < to.x; i++ {
				c[Pos{i, from.y}] = 1
			}
		} else if to.x < from.x {
			for i := to.x + 1; i < from.x; i++ {
				c[Pos{i, from.y}] = 1
			}
		}
		if from.y < to.y {
			for i := from.y + 1; i < to.y; i++ {
				c[Pos{from.x, i}] = 1
			}
		} else if to.y < from.y {
			for i := to.y + 1; i < from.y; i++ {
				c[Pos{from.x, i}] = 1
			}
		}

	}
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n")

	cave := buildCave(input)

	partOne := dropSand(cave, Pos{500, 0})
	fmt.Println(partOne)

	partTwo := dropSand2(cave, Pos{500, 0})
	fmt.Println(partTwo)
}

func buildCave(input []string) Cave {
	c := make(Cave)
	for _, i := range input {
		if i != "" {
			c.addRocks(i)
		}
	}
	return c
}

func dropSand(c Cave, start Pos) int {
	maxY := c.maxY()
	current := start

	for {
		next := Pos{current.x, current.y + 1}
		if c[next] > 0 { // blocked
			next = Pos{current.x - 1, current.y + 1}
			if c[next] > 0 { // still blocked
				next = Pos{current.x + 1, current.y + 1}
				if c[next] > 0 { // can't fall, add to cave
					c[current] = 2
					current = start
					continue
				}
			}
		}
		current = next
		if current.y > maxY {
			break
		}
	}

	return c.countSand()
}

func dropSand2(c Cave, start Pos) int {
	maxY := c.maxY()
	current := start

	for {
		next := Pos{current.x, current.y + 1}
		if c[next] > 0 { // blocked
			next = Pos{current.x - 1, current.y + 1}
			if c[next] > 0 { // still blocked
				next = Pos{current.x + 1, current.y + 1}
				if c[next] > 0 { // can't fall, add to cave
					c[current] = 2
					if current == start {
						break
					}
					current = start
					continue
				}
			}
		}
		if current.y == maxY+1 { // hit the floor
			c[current] = 2
			current = start
		} else {
			current = next
		}
	}

	return c.countSand()
}
