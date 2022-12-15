package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Pos struct {
	x int
	y int
}

type Sensor struct {
	position         Pos
	beacon           Pos
	distanceToBeacon int
}

func (s Sensor) clearCellsInRow(row int) map[Pos]int {
	cells := make(map[Pos]int)
	start := Pos{s.position.x, row}

	offset := s.position.distanceTo(start)

	if offset > 0 {
		distance := s.distanceToBeacon - offset

		for x := start.x - distance; x <= start.x+distance; x++ {
			cell := Pos{x, row}
			if cell != s.beacon {
				cells[Pos{x, row}] = 1
			}
		}
	}

	return cells
}

func (p Pos) distanceTo(other Pos) int {
	x := math.Abs(float64(p.x - other.x))
	y := math.Abs(float64(p.y - other.y))

	return int(x + y)
}

func clearInRow(sensors []Sensor, row int) int {
	cells := make(map[Pos]int)
	for _, s := range sensors {
		c := s.clearCellsInRow(row)
		for k, _ := range c {
			cells[k] = 1
		}
	}
	return len(cells)
}

func findBeacon(sensors []Sensor, max int) Pos {
	result := Pos{0, 0}
	rowSensors := make(map[int][]Sensor)
	for _, s := range sensors {
		startY := s.position.y - s.distanceToBeacon
		endY := s.position.y + s.distanceToBeacon
		if startY < 0 {
			startY = 0
		}
		if endY > max {
			endY = max
		}
		for y := startY; y <= endY; y++ {
			rowSensors[y] = append(rowSensors[y], s)
		}
	}

	found := false
	positions := make(map[Pos]bool)
	row := 0
	for y := 0; y <= max; y++ {
		visited := make(map[Pos]bool)
		for x := 0; x <= max; x++ {
			for _, s := range rowSensors[y] {
				p := Pos{x, y}
				if s.position.distanceTo(p) <= s.distanceToBeacon {
					visited[p] = true
				}
			}
		}
		if len(visited) == max {
			found = true
			row = y
			positions = visited
			break
		}
	}

	if found {
		for x := 0; x <= max; x++ {
			if !positions[Pos{x, row}] {
				result = Pos{x, row}
			}
		}
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	sensors := []Sensor{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("(-*[0-9]+)")
		found := re.FindAllString(line, -1)
		if len(found) == 4 {
			xs, _ := strconv.Atoi(found[0])
			ys, _ := strconv.Atoi(found[1])
			xb, _ := strconv.Atoi(found[2])
			yb, _ := strconv.Atoi(found[3])
			bPos := Pos{xb, yb}
			sPos := Pos{xs, ys}
			s := Sensor{Pos{xs, ys}, bPos, sPos.distanceTo(bPos)}
			sensors = append(sensors, s)
		}
	}

	//y := 2000000
	max := 4000000
	//fmt.Printf("Covers in row %d: %d\n", y, clearInRow(sensors, y))
	fmt.Printf("Beacon at  %v\n", findBeacon(sensors, max))
}
