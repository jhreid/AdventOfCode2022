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

func (d Droplet) fillVoids() {
	maxX := d.maxX()
	maxY := d.maxY()
	maxZ := d.maxZ()
	filledCubes := make(map[Cube]int)
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			for z := 0; z <= maxZ; z++ {
				test := Cube{x, y, z}
				if _, exists := d.cubes[test]; exists {
					filledCubes[test] = 1
				}
			}
		}
	}

	filledCubes = floodFill(filledCubes, maxX, maxY, maxZ)

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			for z := 0; z <= maxZ; z++ {
				test := Cube{x, y, z}
				if filledCubes[test] == 0 {
					filledCubes[test] = 2
				}
			}
		}
	}
	for x := -1; x <= maxX+1; x++ {
		for y := -1; y <= maxY+1; y++ {
			for z := -1; z <= maxZ+1; z++ {
				test := Cube{x, y, z}
				if filledCubes[test] == -1 {
					delete(filledCubes, test)
				}
			}
		}
	}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			for z := 0; z <= maxZ; z++ {
				test := Cube{x, y, z}
				if filledCubes[test] == 2 {
					d.addCube(test)
				}
			}
		}
	}
}

func floodFill(filledCubes map[Cube]int, maxX int, maxY int, maxZ int) map[Cube]int {
	queue := []Cube{{0, 0, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if filledCubes[current] == 0 {
			filledCubes[current] = -1
			if current.x > 0 {
				queue = append(queue, Cube{current.x - 1, current.y, current.z})
			}
			if current.x < maxX {
				queue = append(queue, Cube{current.x + 1, current.y, current.z})
			}
			if current.y > 0 {
				queue = append(queue, Cube{current.x, current.y - 1, current.z})
			}
			if current.y < maxY {
				queue = append(queue, Cube{current.x, current.y + 1, current.z})
			}
			if current.z > 0 {
				queue = append(queue, Cube{current.x, current.y, current.z - 1})
			}
			if current.z < maxZ {
				queue = append(queue, Cube{current.x, current.y, current.z + 1})
			}
		}
	}

	return filledCubes
}

func (d Droplet) maxX() int {
	x := 0
	for c := range d.cubes {
		if c.x > x {
			x = c.x
		}
	}
	return x
}

func (d Droplet) maxY() int {
	y := 0
	for c := range d.cubes {
		if c.y > y {
			y = c.y
		}
	}
	return y
}

func (d Droplet) maxZ() int {
	z := 0
	for c := range d.cubes {
		if c.z > z {
			z = c.z
		}
	}
	return z
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

	droplet.fillVoids()
	fmt.Printf("Part two: %d\n\n", droplet.surfaceArea())
}
