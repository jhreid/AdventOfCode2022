package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Knot struct {
	number    int
	pos       Pos
	history   []Pos
	locations map[Pos]int
}

func (k *Knot) follow(leader *Knot) {
	x := leader.pos.x
	y := leader.pos.y
	dx := x - k.pos.x
	dy := y - k.pos.y

	if dy == 0 { // horizontal
		if dx > 1 {
			k.pos.x += 1
		} else if dx < -1 {
			k.pos.x -= 1
		}
	} else if dx == 0 { // vertical
		if dy > 1 {
			k.pos.y += 1
		} else if dy < -1 {
			k.pos.y -= 1
		}
	} else if (math.Abs(float64(dy)) == 2 && math.Abs(float64(dx)) > 0) || (math.Abs(float64(dx)) == 2 && math.Abs(float64(dx)) >= 1) {
		if dx > 0 {
			k.pos.x += 1
		} else {
			k.pos.x -= 1
		}
		if dy > 0 {
			k.pos.y += 1
		} else {
			k.pos.y -= 1
		}
	}

	k.history = append(k.history, k.pos)
	k.locations[k.pos] += 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	tail := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 9}
	eight := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 8}
	seven := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 7}
	six := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 6}
	five := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 5}
	four := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 4}
	three := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 3}
	two := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 2}
	one := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 1}
	head := Knot{pos: Pos{0, 0}, locations: map[Pos]int{}, history: []Pos{}, number: 0}

	//rope := []*Knot{&head, &one, &two, &three, &four, &five, &six, &seven, &eight, &tail}
	rope := []*Knot{&head, &one, &two, &three, &four, &five, &six, &seven, &eight, &tail}

	for _, k := range rope {
		k.history = append(k.history, k.pos)
		k.locations[k.pos] = 1
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			instruction := strings.Split(line, " ")
			direction := instruction[0]
			distance, _ := strconv.Atoi(instruction[1])

			processInstruction2(rope, distance, direction)
		}
	}

	fmt.Printf("Positions visited: %d\n", len(rope[9].locations))
}

func processInstruction2(rope []*Knot, distance int, direction string) {
	head := rope[0]
	for d := 1; d <= distance; d++ {
		switch direction {
		case "R":
			head.pos.x += 1
		case "L":
			head.pos.x -= 1
		case "U":
			head.pos.y += 1
		case "D":
			head.pos.y -= 1
		}

		rope[1].follow(head)
		for i, k := range rope[2:] {
			k.follow(rope[i+1])
		}
	}
}
