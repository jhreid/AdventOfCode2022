package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (p Pos) distanceTo(other Pos) int {
	x := Abs(p.x - other.x)
	y := Abs(p.y - other.y)

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

	for y := 0; y <= max+1; y++ {
		lines := [][]int{}
		for _, s := range sensors {
			distFromRow := Abs(s.position.y - y)
			offset := s.distanceToBeacon - distFromRow
			if offset >= 0 {
				lines = append(lines, []int{s.position.x - offset, s.position.x + offset})
			}
		}

		sort.Slice(lines, func(i, j int) bool { // sort on first value
			return lines[i][0] < lines[j][0]
		})

		testEnd := lines[0][1] // first line end

		for _, nextRange := range lines[1:] { // iterate over rest looking for overlaps or gaps
			if nextRange[0] <= testEnd { // overlap
				testEnd = Max(nextRange[1], testEnd)
			} else {
				if nextRange[0]-1 > testEnd { // gap
					return Pos{nextRange[0] - 1, y}
				}
				testEnd = nextRange[1]
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
	beacon := findBeacon(sensors, max)
	fmt.Printf("Beacon at  %v\n", beacon)
	fmt.Printf("Signal strength  %d\n", beacon.x*4000000+beacon.y)
}
