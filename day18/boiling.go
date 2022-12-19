package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

type Droplet struct {
	cubes map[Cube]int
}

func (d Droplet) addCube(c Cube) {
	d.cubes[c] = 0
	test := Cube{c.x - 1, c.y, c.z}
	if _, exists := d.cubes[test]; exists {
		d.cubes[test]++
		d.cubes[c]++
	}
	test = Cube{c.x + 1, c.y, c.z}
	if _, exists := d.cubes[test]; exists {
		d.cubes[test]++
		d.cubes[c]++
	}
	test = Cube{c.x, c.y - 1, c.z}
	if _, exists := d.cubes[test]; exists {
		d.cubes[test]++
		d.cubes[c]++
	}
	test = Cube{c.x, c.y + 1, c.z}
	if _, exists := d.cubes[test]; exists {
		d.cubes[test]++
		d.cubes[c]++
	}
	test = Cube{c.x, c.y, c.z - 1}
	if _, exists := d.cubes[test]; exists {
		d.cubes[test]++
		d.cubes[c]++
	}
	test = Cube{c.x, c.y, c.z + 1}
	if _, exists := d.cubes[test]; exists {
		d.cubes[test]++
		d.cubes[c]++
	}
}

func (d Droplet) surfaceArea() int {
	area := len(d.cubes) * 6
	hiddenSurfaces := 0
	for _, v := range d.cubes {
		hiddenSurfaces += v
	}
	return area - hiddenSurfaces
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	droplet := Droplet{make(map[Cube]int)}
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		droplet.addCube(Cube{x, y, z})
	}

	fmt.Printf("Part one: %d\n\n", droplet.surfaceArea())
}
